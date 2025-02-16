package domain

import (
	"time"

	"github.com/hoshitocat/upsider-coding-test/internal/timex"
	"github.com/oklog/ulid/v2"
)

const (
	InvoiceStatusUnprocessed = "unprocessed" // 未処理
	InvoiceStatusProcessing  = "processing"  // 処理中
	InvoiceStatusPaid        = "paid"        // 支払い済み
	InvoiceStatusError       = "error"       // エラー
)

type Invoice struct {
	ID                string     `json:"id"`
	CompanyID         string     `json:"company_id"`
	BusinessPartnerID string     `json:"business_partner_id"`
	IssueDate         timex.Date `json:"issue_date"`
	PaymentAmount     float64    `json:"payment_amount"`
	FeeRate           float64    `json:"fee_rate"`
	FeeAmount         float64    `json:"fee_amount"`
	TaxRate           float64    `json:"tax_rate"`
	TaxAmount         float64    `json:"tax_amount"`
	TotalAmount       float64    `json:"total_amount"`
	DueDate           timex.Date `json:"due_date"`
	Status            string     `json:"status"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

func NewInvoice(
	companyID string,
	businessPartnerID string,
	issueDate timex.Date,
	paymentAmount float64,
	dueDate timex.Date,
) *Invoice {
	// TODO: 外部から渡せるようにしたほうが良さそうだが、一旦要件では固定値で良さそうなのでこちらに定義
	const (
		feeRate = 0.04
		taxRate = 0.1
	)

	invoice := &Invoice{
		ID:                ulid.Make().String(), // TODO: IDの採番についてはチームで要相談
		CompanyID:         companyID,
		BusinessPartnerID: businessPartnerID,
		IssueDate:         issueDate,
		PaymentAmount:     paymentAmount,
		FeeRate:           feeRate,
		TaxRate:           taxRate,
		DueDate:           dueDate,
		Status:            InvoiceStatusUnprocessed,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
	invoice.FeeAmount = invoice.calculateFeeAmount()
	invoice.TaxAmount = invoice.calculateTaxAmount()
	invoice.TotalAmount = invoice.calculateTotalAmount()

	// TODO: validate invoice

	return invoice
}

func (i *Invoice) calculateFeeAmount() float64 {
	return i.PaymentAmount * i.FeeRate
}

func (i *Invoice) calculateTaxAmount() float64 {
	return i.FeeAmount * i.TaxRate
}

func (i *Invoice) calculateTotalAmount() float64 {
	return i.PaymentAmount + i.FeeAmount + i.TaxAmount
}
