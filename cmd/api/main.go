package main

import (
	"github.com/Guiziin227/goCleanArc/internal/handlers"
	"github.com/Guiziin227/goCleanArc/internal/repositories"
	"github.com/Guiziin227/goCleanArc/internal/usecases"
)

func main() {

	repo := repositories.NewRepositories()
	useCase := usecases.NewUseCases(repo)

	h := handlers.NewHandlers(useCase)

	err := h.Listen(8080)
	if err != nil {
		return
	}
}
