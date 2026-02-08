package ports

import (
	"context"

	"github.com/go-hexagonal-practice/internal/core/domain/user"
)

type UserRepositoryPort interface {
	GetByEmail(ctx context.Context, email string) (*user.User, *user.UserCredentials)
	// UpdateLastLogin
}
