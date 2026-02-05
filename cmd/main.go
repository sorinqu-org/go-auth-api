package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/sorinqu-org/go-auth-api/cmd/api"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, os.Getenv("GOOSE_DBSTRING"))
	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
		return
	}

	server := api.NewAPIServer(":8080", conn)
	if err := server.Run(); err != nil {
		slog.Error("Failed to run server", "error", err)
	}
}
