package repositories

import (
	"context"
	"database/sql"

	"github.com/Guiziin227/goCleanArc/internal/models"
	"github.com/Guiziin227/goCleanArc/internal/repositories/users"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetAll(ctx context.Context) ([]models.User, error)
	Add(ctx context.Context, newUser models.User) error
	EmailExists(ctx context.Context, email string) (bool, error)
	GetById(ctx context.Context, id uuid.UUID) (models.User, error)
}

type Repositories struct {
	User UserRepository
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		User: users.NewUsers(db),
	}
}
