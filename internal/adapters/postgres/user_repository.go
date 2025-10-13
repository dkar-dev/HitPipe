// services/User/internal/adapters/postgres/user_repository.go
package postgres

import (
	"Main/internal/domain"
	"context"
	"github.com/jmoiron/sqlx"
)

// UserRepository реализует интерфейс ports.UserRepository для Postgres.
type UserRepository struct {
	db *sqlx.DB
}

// NewUserRepository - это конструктор для нашего репозитория.
func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Save реализует метод сохранения пользователя.
func (r *UserRepository) Save(ctx context.Context, user *domain.User) error {

	query := `INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id, created_at`

	err := r.db.QueryRowxContext(ctx, query, user.Email, user.Password).
		Scan(&user.ID, &user.CreatedAt)

	return err
}
