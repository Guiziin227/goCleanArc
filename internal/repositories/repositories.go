package repositories

import (
	"github.com/Guiziin227/goCleanArc/internal/models"
	"github.com/Guiziin227/goCleanArc/internal/repositories/users"
	"github.com/google/uuid"
)

type Repositories struct {
	User interface {
		GetAll() []models.User
		Add(newUser models.User)
		EmailExists(email string) bool
		GetById(id uuid.UUID) models.User
	}
}

func NewRepositories() *Repositories {
	return &Repositories{
		User: users.NewUsers(),
	}
}
