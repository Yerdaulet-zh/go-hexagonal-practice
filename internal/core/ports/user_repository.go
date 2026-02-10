package ports

import (
	"context"

	domain_profile "github.com/go-hexagonal-practice/internal/core/domain/profile"
	domain_sessions "github.com/go-hexagonal-practice/internal/core/domain/sessions"
	domain_user "github.com/go-hexagonal-practice/internal/core/domain/user"
)

type UserPort interface {
	GetUserByEmail(ctx context.Context, email string) (*domain_user.User, error)
	CreateUser(ctx context.Context, user *domain_user.User, userCredentials *domain_user.UserCredentials, userSession *domain_sessions.UserSessions, userProfileRecord *domain_profile.UserProfile) (*domain_sessions.UserSessions, error)
}
