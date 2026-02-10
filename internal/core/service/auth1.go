package service

// import (
// 	"context"
// 	"errors"
// 	"fmt"

// 	"github.com/google/uuid"
// 	"golang.org/x/crypto/bcrypt"

// 	domain_sessions "github.com/go-hexagonal-practice/internal/core/domain/sessions"
// 	app_ports "github.com/go-hexagonal-practice/internal/core/ports"
// )

// type authService struct {
// 	userRepo    app_ports.UserRepositoryPort
// 	sessionRepo app_ports.SessionRepositoryPort
// }

// func NewAuthService(u app_ports.UserRepositoryPort, s app_ports.SessionRepositoryPort) app_ports.AuthServicePort {
// 	return &authService{
// 		userRepo:    u,
// 		sessionRepo: s,
// 	}
// }

// func (s *authService) Login(ctx context.Context, email, password string, meta app_ports.SessionMetadata) (*domain_sessions.UserSessions, string, error) {
// 	// 1. Retrieve User & Credentials
// 	u, creds, err := s.userRepo.GetByEmail(ctx, email)
// 	if err != nil {
// 		return nil, "", errors.New("authentication failed: invalid credentials")
// 	}

// 	// 2. Business Rule: Check User Status
// 	if u.UserStatus != "active" {
// 		return nil, "", fmt.Errorf("authentication failed: account status is %s", u.UserStatus)
// 	}

// 	// 3. Verify Password
// 	err = bcrypt.CompareHashAndPassword([]byte(creds.PasswordHash), []byte(password))
// 	if err != nil {
// 		return nil, "", errors.New("authentication failed: invalid credentials")
// 	}

// 	// 4. Generate Session Logic
// 	// In a real app, use a JWT library or secure random string generator here
// 	tokenString := uuid.New().String()

// 	newSession := &domain_sessions.UserSessions{
// 		ID:               uuid.New(), // Generate the primary key for the session
// 		UserID:           u.ID,
// 		RefreshTokenHash: hashToken(tokenString),
// 		IPAddress:        meta.IPAddress,
// 		UserAgent:        &meta.UserAgent,
// 	}

// 	// 5. Persist Session
// 	if err := s.sessionRepo.Create(ctx, newSession); err != nil {
// 		return nil, "", fmt.Errorf("failed to create session: %w", err)
// 	}

// 	return newSession, tokenString, nil
// }

// func (s *authService) Logout(ctx context.Context, sessionID string) error {
// 	id, err := uuid.Parse(sessionID)
// 	if err != nil {
// 		return errors.New("invalid session id format")
// 	}

// 	return s.sessionRepo.Revoke(ctx, id)
// }

// func hashToken(token string) string {
// 	// Recommendation: Use SHA256 for the stored hash of the refresh token
// 	return token
// }
