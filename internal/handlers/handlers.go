package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Guiziin227/goCleanArc/internal/usecases"
)

type Handlers struct {
	usecases *usecases.UseCases
}

func NewHandlers(usecases *usecases.UseCases) *Handlers {
	return &Handlers{
		usecases: usecases,
	}
}

func (h *Handlers) Listen(port int) error {
	h.registerUserEndpoints()

	slog.Info("Listening on", "port", port)

	return http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
