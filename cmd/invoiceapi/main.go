package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/hoshitocat/upsider-coding-test/cmd/invoiceapi/internal/config"
)

func main() {
	ctx := context.Background()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed to create config: %v\n", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/invoices", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("請求書が作成されました"))
	})
	mux.HandleFunc("GET /api/invoices", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("請求書一覧が取得されました"))
	})

	server := &http.Server{
		Addr:    cfg.Port,
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
