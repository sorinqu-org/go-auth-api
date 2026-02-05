package user

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router chi.Router) {
	router.Route("/users", func(r chi.Router) {
		r.Post("/login", h.Login)
		r.Post("/register", h.Register)
	})
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
}
