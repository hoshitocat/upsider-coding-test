package domain

import (
	"context"

	"github.com/hoshitocat/upsider-coding-test/internal/timex"
)

type Repositories struct {
	InvoiceRepository InvoiceRepository
}

type InvoiceRepository interface {
	CreateInvoice(ctx context.Context, invoice *Invoice) error
	ListInvoices(ctx context.Context, beginDate, endDate timex.Date) ([]*Invoice, error)
}
