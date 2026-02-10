package ports

import (
	"context"

	domain_sessions "github.com/go-hexagonal-practice/internal/core/domain/sessions"
)

type UserUseCase interface {
	Register(ctx context.Context, email string, password string, ipAddress string, userAgent *string) (*domain_sessions.UserSessions, error)
}
