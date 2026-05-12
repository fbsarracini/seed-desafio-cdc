package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/fbsarracini/seed-desafio-cdc/internal/author"
	"github.com/fbsarracini/seed-desafio-cdc/internal/config"
	"github.com/fbsarracini/seed-desafio-cdc/internal/database"
)

func main() {
	// logging estruturado (JSON)
	// slog.NewJSONHandler escreve logs em JSON
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// carrega configuração
	cfg := config.Load()

	// abri conexão com o banco de dados
	db, err := database.Open(cfg.DatabaseURL)
	if err != nil {
		slog.Error("falha ao iniciar banco de dados", "err", err)
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	defer db.Close()

	// cria roteador HTTP
	mux := http.NewServeMux()

	// registra rota POST /authors
	// author.CreateHandler(db) é a FACTORY FUNCTION que cria o handler
	// retorna uma http.HandlerFunc
	mux.HandleFunc("POST /authors", author.CreateHandler(db))

	// Configura o servidor HTTP com timeouts
	srv := &http.Server{
		Addr:         cfg.Addr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	slog.Info("iniciando servidor", "addr", cfg.Addr)

	// ListenAndServe
	if err := srv.ListenAndServe(); err != nil {
		slog.Error("server error", "err", err)
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
