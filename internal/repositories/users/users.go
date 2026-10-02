package users

import (
	"sync"

	"github.com/Guiziin227/goCleanArc/internal/models"
)

type Users struct {
	mu    sync.RWMutex
	users []models.User
}

func NewUsers() *Users {
	return &Users{
		users: make([]models.User, 0),
	}
}

func (u *Users) GetAll() []models.User {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return u.users
}

func (u *Users) EmailExists(email string) bool {
	u.mu.RLock()
	defer u.mu.RUnlock()

	for _, user := range u.users {
		if user.Email == email {
			return true
		}
	}
	return false
}

func (u *Users) Add(newUser models.User) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.users = append(u.users, newUser)
}
