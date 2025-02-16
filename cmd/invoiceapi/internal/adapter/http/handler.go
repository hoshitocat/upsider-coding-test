package http

import "github.com/hoshitocat/upsider-coding-test/cmd/invoiceapi/internal/usecase"

type Handlers struct {
	InvoiceHandler *invoiceHandler
}

func NewHandlers(interactors *usecase.Interactors) *Handlers {
	return &Handlers{
		InvoiceHandler: newInvoiceHandler(interactors),
	}
}
