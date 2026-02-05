package api

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sorinqu-org/go-auth-api/service/user"
)

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := chi.NewRouter()
	apiRouter := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(apiRouter)

	router.Mount("/api/v1", apiRouter)

	slog.Info("Server has started", "addres", s.addr)
	return http.ListenAndServe(s.addr, router)
}
