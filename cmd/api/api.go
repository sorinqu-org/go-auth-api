package api

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := chi.NewRouter()
	router.Route("/api/v1", func(r chi.Router) {

	})
	return http.ListenAndServe(s.addr, router)
}
