// services/User/internal/adapters/postgres/user_repository.go
package postgres

import (
	"github.com/dkar-dev/HitPipe/internal/domain"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func (r *UserRepository) Save(u *domain.User) (string, error) {

	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`

	_, err := r.db.Exec(query, u.Username, u.Email, u.Password)
	if err != nil {

	}

	return "no", nil
}
