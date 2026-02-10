package ports

import (
	"context"

	app_sessions "github.com/go-hexagonal-practice/internal/core/domain/sessions"
	"github.com/google/uuid"
)

type SessionRepositoryPort interface {
	Create(ctx context.Context, session *app_sessions.UserSessions) error

	FindByID(ctx context.Context, sessionID uuid.UUID) (*app_sessions.UserSessions, error)

	Revoke(ctx context.Context, sessionID uuid.UUID) error
}
