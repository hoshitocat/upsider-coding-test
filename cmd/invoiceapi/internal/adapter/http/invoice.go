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

	w.WriteHeader(http.StatusCreated)
}
