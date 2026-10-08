package usecases

import (
	"context"
	"fmt"

	"github.com/Guiziin227/goCleanArc/internal/models"
	"github.com/Guiziin227/goCleanArc/internal/repositories"
	"github.com/google/uuid"
)

type UseCases struct {
	Repos *repositories.Repositories
}

func NewUseCases(repos *repositories.Repositories) *UseCases {
	return &UseCases{
		Repos: repos,
	}
}

func (u UseCases) GetAll(ctx context.Context) ([]models.User, error) {
	return u.Repos.User.GetAll(ctx)
}

func (u UseCases) GetById(ctx context.Context, id uuid.UUID) (models.User, error) {
	user, err := u.Repos.User.GetById(ctx, id)
	if err != nil {
		return models.User{}, err
	}
	if user.ID == uuid.Nil {
		return models.User{}, fmt.Errorf("user not found")
	}
	return user, nil
}

func (u UseCases) Add(ctx context.Context, newUser models.CreateUserRequest) (uuid.UUID, error) {

	exists, err := u.Repos.User.EmailExists(ctx, newUser.Email)
	if err != nil {
		return uuid.Nil, err
	}
	if exists {
		return uuid.Nil, fmt.Errorf("email already exists")
	}

	repoReq := models.User{
		ID:    uuid.New(),
		Name:  newUser.Name,
		Email: newUser.Email,
	}

	if err := u.Repos.User.Add(ctx, repoReq); err != nil {
		return uuid.Nil, err
	}

	return repoReq.ID, nil
}
