package repositories

import (
	"github.com/Guiziin227/goCleanArc/internal/models"
	"github.com/Guiziin227/goCleanArc/internal/repositories/users"
)

type Repositories struct {
	User interface {
		GetAll() []models.User
		Add(newUser models.User)
		EmailExists(email string) bool
	}
}

func NewRepositories() *Repositories {
	return &Repositories{
		User: users.NewUsers(),
	}
}
