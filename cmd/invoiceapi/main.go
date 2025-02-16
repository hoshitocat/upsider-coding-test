package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	adapter "github.com/hoshitocat/upsider-coding-test/cmd/invoiceapi/internal/adapter/http"
	"github.com/hoshitocat/upsider-coding-test/cmd/invoiceapi/internal/config"
	"github.com/hoshitocat/upsider-coding-test/cmd/invoiceapi/internal/domain"
	"github.com/hoshitocat/upsider-coding-test/cmd/invoiceapi/internal/infra/db"
	"github.com/hoshitocat/upsider-coding-test/cmd/invoiceapi/internal/usecase"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed to create config: %v\n", err)
	}

	invoiceDB, err := sqlx.Open("mysql", cfg.Database.DSN())
	if err != nil {
		log.Fatalf("Failed to open db: %v\n", err)
	}
	defer invoiceDB.Close()
	err = invoiceDB.PingContext(ctx)
	if err != nil {
		log.Fatalf("Failed to ping db: %v\n", err)
	}

	repos := domain.Repositories{}
	db.InitRepositories(invoiceDB, &repos)
	interactors := usecase.NewInteractors(repos)
	handlers := adapter.NewHandlers(interactors)

	// TODO: 認証基盤やLog基盤を導入する必要があります
	//       認証用のミドルウェアを作成し、リクエストヘッダーから認証トークンを取得し、
	//       トークンの検証を行い、認証が通ったユーザーのみがAPIを利用できるようにする必要があります
	//       また、ログ基盤を導入し、APIのリクエストとレスポンスをログに残したり、OpenTelemetryに準拠したTraceIDやSpanIDの発行ができるような実装を導入したい

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/invoices", handlers.InvoiceHandler.CreateInvoice)
	mux.HandleFunc("GET /api/invoices", handlers.InvoiceHandler.ListInvoices)

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: mux,
	}

	idleConnCh := make(chan struct{})
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt)
		signal.Notify(sigCh, syscall.SIGTERM)

		<-sigCh

		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("Failed to shutdown: %v\n", err)
		}

		close(sigCh)
		close(idleConnCh)
	}()

	log.Printf("Server is ready to handle requests at %s", server.Addr)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v\n", err)
	}

	<-idleConnCh

	log.Printf("Server closed")
}
