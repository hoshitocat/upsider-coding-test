package usecase

import "github.com/hoshitocat/upsider-coding-test/cmd/invoiceapi/internal/domain"

type Interactors struct {
	InvoiceInteractor *InvoiceInteractor
}

func NewInteractors(repos domain.Repositories) *Interactors {
	return &Interactors{
		InvoiceInteractor: newInvoiceInteractor(repos.InvoiceRepository),
	}
}
