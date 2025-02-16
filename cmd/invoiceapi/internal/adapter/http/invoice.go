package http

import (
	"encoding/json"
	"net/http"

	"github.com/hoshitocat/upsider-coding-test/cmd/invoiceapi/internal/domain"
	"github.com/hoshitocat/upsider-coding-test/cmd/invoiceapi/internal/usecase"
	"github.com/hoshitocat/upsider-coding-test/internal/timex"
)

type invoiceHandler struct {
	invoiceInteractor *usecase.InvoiceInteractor
}

func newInvoiceHandler(interactors *usecase.Interactors) *invoiceHandler {
	return &invoiceHandler{invoiceInteractor: interactors.InvoiceInteractor}
}

type createInvoiceRequest struct {
	CompanyID         string     `json:"company_id"`
	BusinessPartnerID string     `json:"business_partner_id"`
	IssueDate         timex.Date `json:"issue_date"`
	DueDate           timex.Date `json:"due_date"`
	PaymentAmount     float64    `json:"payment_amount"`
}

func (h *invoiceHandler) CreateInvoice(w http.ResponseWriter, r *http.Request) {
	var req createInvoiceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: 請求リクエストのバリデーション

	invoice := domain.NewInvoice(req.CompanyID, req.BusinessPartnerID, req.IssueDate, req.PaymentAmount, req.DueDate)
	if err := h.invoiceInteractor.CreateInvoice(r.Context(), invoice); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(invoice)
	// TODO: エラーハンドリング
}

func (h *invoiceHandler) ListInvoices(w http.ResponseWriter, r *http.Request) {
	begin := r.URL.Query().Get("begin")
	if begin == "" {
		http.Error(w, "begin is required", http.StatusBadRequest)
		return
	}

	beginDate, err := timex.NewDateFromString(begin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	end := r.URL.Query().Get("end")
	if end == "" {
		http.Error(w, "end is required", http.StatusBadRequest)
		return
	}

	endDate, err := timex.NewDateFromString(end)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	invoices, err := h.invoiceInteractor.ListInvoices(r.Context(), beginDate, endDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(invoices) == 0 {
		invoices = []*domain.Invoice{}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string][]*domain.Invoice{"invoices": invoices})
	// TODO: エラーハンドリング
}
