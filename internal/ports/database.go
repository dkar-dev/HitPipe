// services/User/internal/ports/database.go
package ports

import (
	"context"
	"github.com/dkar-dev/HitPipe/internal/domain"
)

type UserRepository interface {
	Save(ctx context.Context, user *domain.User) error
}
