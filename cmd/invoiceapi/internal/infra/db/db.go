package db

import (
	"github.com/hoshitocat/upsider-coding-test/cmd/invoiceapi/internal/domain"
	"github.com/jmoiron/sqlx"
)

func InitRepositories(db *sqlx.DB, repos *domain.Repositories) {
	repos.InvoiceRepository = newInvoiceRepository(db)
}
