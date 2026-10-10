package users

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Guiziin227/goCleanArc/internal/models"
	"github.com/google/uuid"
)

type Users struct {
	db *sql.DB
}

func NewUsers(db *sql.DB) *Users {
	return &Users{
		db: db,
	}
}

func (u *Users) GetAll(ctx context.Context) ([]models.User, error) {
	rows, err := u.db.QueryContext(ctx, ""+
		"SELECT id, name, email "+
		"FROM users "+
		"ORDER BY name")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]models.User, 0)

	for rows.Next() {
		var user models.User

		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email); err != nil {
			return nil, err
		}

		result = append(result, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (u *Users) GetById(ctx context.Context, id uuid.UUID) (models.User, error) {

	var user models.User

	err := u.db.QueryRowContext(ctx, "SELECT id, name, email "+
		"FROM users "+
		"WHERE id = $1", id).Scan(
		&user.ID,
		&user.Name,
		&user.Email)

	if errors.Is(err, sql.ErrNoRows) {
		return models.User{}, nil
	}

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (u *Users) Add(ctx context.Context, newUser models.User) error {

	_, err := u.db.ExecContext(ctx,
		"INSERT INTO users (id, name, email) "+
			"VALUES ($1, $2, $3)",
		newUser.ID,
		newUser.Name,
		newUser.Email)

	if err != nil {
		return err
	}

	return nil
}

func (u *Users) EmailExists(ctx context.Context, email string) (bool, error) {
	var exists bool
	err := u.db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", email).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (u *Users) DeleteById(ctx context.Context, id uuid.UUID) error {
	result, err := u.db.ExecContext(ctx, "DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}
