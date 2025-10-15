// services/User/internal/domain/types.go
package domain

import "time"

// User - это наша чистая доменная модель.
// Она ничего не знает о базе данных или gRPC.
type User struct {
	ID        string    `db:"id"`
	Username  string    `db:"username"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Provider  string    `db:"provider"`
	CreatedAt time.Time `db:"created_at"`
}
