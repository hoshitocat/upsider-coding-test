package domain

import (
	"context"
)

type Repositories struct {
	InvoiceRepository InvoiceRepository
}

type InvoiceRepository interface {
	CreateInvoice(ctx context.Context, invoice *Invoice) error
}
