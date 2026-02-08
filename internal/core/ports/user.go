package ports

import (
	"context"

	"github.com/go-hexagonal-practice/internal/core/domain/sessions"
	"github.com/go-hexagonal-practice/internal/core/domain/user"
	"github.com/google/uuid"
)

type UserRepositoryPort interface {
	GetByEmail(ctx context.Context, email string) (*user.User, *user.UserCredentials, error)
	// UpdateLastLogin
}

type SessionRepositoryPort interface {
	Create(ctx context.Context, session *sessions.UserSessions) error

	FindByID(ctx context.Context, sessionID uuid.UUID) (*sessions.UserSessions, error)

	Revoke(ctx context.Context, sessionID uuid.UUID) error
}
