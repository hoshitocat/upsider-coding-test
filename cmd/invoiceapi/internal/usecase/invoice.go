package usecase

import (
	"context"
	"fmt"

	"github.com/hoshitocat/upsider-coding-test/cmd/invoiceapi/internal/domain"
	"github.com/hoshitocat/upsider-coding-test/internal/timex"
)

type InvoiceInteractor struct {
	invoiceRepo domain.InvoiceRepository
}

func newInvoiceInteractor(invoiceRepo domain.InvoiceRepository) *InvoiceInteractor {
	return &InvoiceInteractor{invoiceRepo: invoiceRepo}
}

func (i *InvoiceInteractor) CreateInvoice(ctx context.Context, invoice *domain.Invoice) error {
	if err := i.invoiceRepo.CreateInvoice(ctx, invoice); err != nil {
		return fmt.Errorf("failed to create invoice: %w", err)
	}

	return nil
}

func (i *InvoiceInteractor) ListInvoices(ctx context.Context, beginDate, endDate timex.Date) ([]*domain.Invoice, error) {
	invoices, err := i.invoiceRepo.ListInvoices(ctx, beginDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to list invoices: %w", err)
	}
	return invoices, nil
}
