package main

import (
	"context"
	"log"
	"net/http"

	"github.com/Guiziin227/goCleanArc/internal/database"
	"github.com/Guiziin227/goCleanArc/internal/handlers"
	"github.com/Guiziin227/goCleanArc/internal/repositories"
	"github.com/Guiziin227/goCleanArc/internal/usecases"
)

func main() {
	db, err := database.InitPostgres(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repositories.NewRepositories(db)
	useCase := usecases.NewUseCases(repo)

	h := handlers.NewHandlers(useCase)

	log.Println("Listening on :8000")

	if err := http.ListenAndServe(":8000", h.Router()); err != nil {
		log.Fatal(err)
	}
}
