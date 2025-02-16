package usecase

import (
	"context"
	"fmt"

	"github.com/hoshitocat/upsider-coding-test/cmd/invoiceapi/internal/domain"
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
