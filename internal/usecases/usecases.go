package usecases

import (
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

func (u UseCases) GetAll() []models.User {
	users := u.Repos.User.GetAll()

	return users
}

func (u UseCases) GetById(id uuid.UUID) (models.User, error) {
	user := u.Repos.User.GetById(id)
	if user.ID == uuid.Nil {
		return models.User{}, fmt.Errorf("user not found")
	}
	return user, nil
}

func (u UseCases) Add(newUser models.CreateUserRequest) (uuid.UUID, error) {

	exists := u.Repos.User.EmailExists(newUser.Email)
	if exists {
		return uuid.Nil, fmt.Errorf("email already exists")
	}

	repoReq := models.User{
		ID:    uuid.New(),
		Name:  newUser.Name,
		Email: newUser.Email,
	}

	u.Repos.User.Add(repoReq)

	return repoReq.ID, nil
}
