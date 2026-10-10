package handlers

import (
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
