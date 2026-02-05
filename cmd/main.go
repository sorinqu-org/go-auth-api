package main

import (
	"log/slog"
	"os"

	"github.com/sorinqu-org/go-auth-api/cmd/api"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		slog.Error("Failed to run server", "error", err)
	}
}
