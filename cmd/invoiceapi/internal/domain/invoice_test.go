package domain_test

import (
	"testing"

	"github.com/hoshitocat/upsider-coding-test/cmd/invoiceapi/internal/domain"
)

func TestInvoice_calculateFeeAmount(t *testing.T) {
	invoice := &domain.Invoice{
		PaymentAmount: 1000,
		FeeRate:       0.1,
	}

	got := domain.ExportedInvoice_calculateFeeAmount(invoice)
	var want float64 = 100

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestInvoice_calculateTaxAmount(t *testing.T) {
	invoice := &domain.Invoice{
		FeeAmount: 100,
		TaxRate:   0.1,
	}

	got := domain.ExportedInvoice_calculateTaxAmount(invoice)
	var want float64 = 10

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestInvoice_calculateTotalAmount(t *testing.T) {
	invoice := &domain.Invoice{
		PaymentAmount: 1000,
		FeeAmount:     100,
		TaxAmount:     10,
	}

	got := domain.ExportedInvoice_calculateTotalAmount(invoice)
	var want float64 = 1110

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
