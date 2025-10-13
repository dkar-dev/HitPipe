// services/User/internal/ports/database.go
package ports

import (
	"Main/internal/domain"
	"context"
)

// UserRepository - это наш порт. Он определяет,
// какие методы для работы с пользователями должно предоставлять хранилище.
type UserRepository interface {
	Save(ctx context.Context, user *domain.User) error
}
