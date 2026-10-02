package handlers

import (
	"fmt"
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
	return http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
