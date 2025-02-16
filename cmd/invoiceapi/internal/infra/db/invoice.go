package db

import (
	"context"
	"fmt"
	"time"

	"github.com/hoshitocat/upsider-coding-test/cmd/invoiceapi/internal/domain"
	"github.com/hoshitocat/upsider-coding-test/internal/timex"
	"github.com/jmoiron/sqlx"
)

type invoiceRepository struct {
	db *sqlx.DB
}

func newInvoiceRepository(db *sqlx.DB) domain.InvoiceRepository {
	return &invoiceRepository{db: db}
}

func (r *invoiceRepository) CreateInvoice(ctx context.Context, do *domain.Invoice) error {
	po, err := do2poInvoice(do)
	if err != nil {
		return fmt.Errorf("failed to convert invoice to po invoice: %w", err)
	}

	query := `	
		INSERT INTO invoices (
			id,
			company_id,
			business_partner_id,
			issue_date,
			payment_amount,
			fee_rate,
			fee_amount,
			tax_rate,
			tax_amount,
			total_amount,
			due_date,
			status_id,
			created_at,
			updated_at
		)
		VALUES (
			:id,
			:company_id,
			:business_partner_id,
			:issue_date,
			:payment_amount,
			:fee_rate,
			:fee_amount,
			:tax_rate,
			:tax_amount,
			:total_amount,
			:due_date,
			:status_id,
			:created_at,
			:updated_at
		);
	`

	_, err = r.db.NamedExecContext(ctx, query, po)
	if err != nil {
		return fmt.Errorf("failed to create invoice: %w", err)
	}

	return nil
}

func (r *invoiceRepository) ListInvoices(ctx context.Context, beginDate, endDate timex.Date) ([]*domain.Invoice, error) {
	query := `
		SELECT * FROM invoices
		WHERE due_date BETWEEN :begin_date AND :end_date
		ORDER BY due_date ASC
	`

	query, args, err := sqlx.Named(query, map[string]any{
		"begin_date": beginDate.String(),
		"end_date":   endDate.String(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to prepare query: %w", err)
	}

	var pos []*invoice
	if err := r.db.SelectContext(ctx, &pos, query, args...); err != nil {
		return nil, fmt.Errorf("failed to list invoices: %w", err)
	}

	dos := make([]*domain.Invoice, len(pos))
	for i, po := range pos {
		dos[i], err = po2doInvoice(po)
		if err != nil {
			return nil, fmt.Errorf("failed to convert invoice to do invoice: %w", err)
		}
	}

	return dos, nil
}

type invoice struct {
	ID                string    `db:"id"`
	CompanyID         string    `db:"company_id"`
	BusinessPartnerID string    `db:"business_partner_id"`
	IssueDate         time.Time `db:"issue_date"`
	PaymentAmount     float64   `db:"payment_amount"`
	FeeRate           float64   `db:"fee_rate"`
	FeeAmount         float64   `db:"fee_amount"`
	TaxRate           float64   `db:"tax_rate"`
	TaxAmount         float64   `db:"tax_amount"`
	TotalAmount       float64   `db:"total_amount"`
	DueDate           time.Time `db:"due_date"`
	StatusID          string    `db:"status_id"`
	CreatedAt         time.Time `db:"created_at"`
	UpdatedAt         time.Time `db:"updated_at"`
}

func do2poStatus(s string) (string, error) {
	/*
		invoice_statuseのシードから生成されたIDを返す
		01JM69M7A9V4Z42DBRFYBXXVKR: 未処理
		01JM69M7AAZ3VBAF2QT9J0Z78P: 処理中
		01JM69M7AAPY0FAC0YDW27WY0C: 支払い済み
		01JM69M7ABS0YJDE2TQ7SXWC47: エラー
	*/
	switch s {
	case domain.InvoiceStatusUnprocessed:
		return "01JM69M7A9V4Z42DBRFYBXXVKR", nil
	case domain.InvoiceStatusProcessing:
		return "01JM69M7AAZ3VBAF2QT9J0Z78P", nil
	case domain.InvoiceStatusPaid:
		return "01JM69M7AAPY0FAC0YDW27WY0C", nil
	case domain.InvoiceStatusError:
		return "01JM69M7ABS0YJDE2TQ7SXWC47", nil
	}

	return "", fmt.Errorf("invalid status: %s", s)
}

func po2doStatus(s string) (string, error) {
	switch s {
	case "01JM69M7A9V4Z42DBRFYBXXVKR":
		return domain.InvoiceStatusUnprocessed, nil
	case "01JM69M7AAZ3VBAF2QT9J0Z78P":
		return domain.InvoiceStatusProcessing, nil
	case "01JM69M7AAPY0FAC0YDW27WY0C":
		return domain.InvoiceStatusPaid, nil
	case "01JM69M7ABS0YJDE2TQ7SXWC47":
		return domain.InvoiceStatusError, nil
	}

	return "", fmt.Errorf("invalid status: %s", s)
}

func do2poInvoice(i *domain.Invoice) (*invoice, error) {
	statusID, err := do2poStatus(i.Status)
	if err != nil {
		return nil, fmt.Errorf("failed to convert status to po status: %w", err)
	}

	return &invoice{
		ID:                i.ID,
		CompanyID:         i.CompanyID,
		BusinessPartnerID: i.BusinessPartnerID,
		IssueDate:         i.IssueDate.Time(),
		PaymentAmount:     i.PaymentAmount,
		FeeRate:           i.FeeRate,
		FeeAmount:         i.FeeAmount,
		TaxRate:           i.TaxRate,
		TaxAmount:         i.TaxAmount,
		TotalAmount:       i.TotalAmount,
		DueDate:           i.DueDate.Time(),
		StatusID:          statusID,
		CreatedAt:         i.CreatedAt,
		UpdatedAt:         i.UpdatedAt,
	}, nil
}

func po2doInvoice(po *invoice) (*domain.Invoice, error) {
	status, err := po2doStatus(po.StatusID)
	if err != nil {
		return nil, fmt.Errorf("failed to convert status to do status: %w", err)
	}

	return &domain.Invoice{
		ID:                po.ID,
		CompanyID:         po.CompanyID,
		BusinessPartnerID: po.BusinessPartnerID,
		IssueDate:         timex.NewDateFromTime(po.IssueDate),
		PaymentAmount:     po.PaymentAmount,
		FeeRate:           po.FeeRate,
		FeeAmount:         po.FeeAmount,
		TaxRate:           po.TaxRate,
		TaxAmount:         po.TaxAmount,
		TotalAmount:       po.TotalAmount,
		DueDate:           timex.NewDateFromTime(po.DueDate),
		Status:            status,
		CreatedAt:         po.CreatedAt,
		UpdatedAt:         po.UpdatedAt,
	}, nil
}
