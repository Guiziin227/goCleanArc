package handlers

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (h *Handlers) Router() http.Handler {

	r := chi.NewRouter() //Cria um novo roteador.

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	r.Route("/users", func(r chi.Router) {
		r.Get("/", h.getAllUsers)
		r.Post("/", h.addUser)
		r.Get("/{id}", h.getUserByID)
		r.Delete("/{id}", h.deleteByID)
	})

	return r

}
