package ports

import (
	"context"

	domain_sessions "github.com/go-hexagonal-practice/internal/core/domain/sessions"
)

type UserUseCase interface {
	Register(ctx context.Context, params RegisterParams) (*domain_sessions.UserSessions, error)
}

type RegisterParams struct {
	Email         string
	Password      string
	FirstName     string
	LastName      *string
	CountryCode   *string
	CountrySource *string
	IPAddress     string
	UserAgent     *string
	Device        *string
}
