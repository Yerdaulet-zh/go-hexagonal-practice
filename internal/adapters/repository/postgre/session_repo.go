package postgre

import (
	"context"
	"errors"

	gorm_sessions "github.com/go-hexagonal-practice/internal/adapters/repository/postgre/persistency/sessions"
	domain_sessions "github.com/go-hexagonal-practice/internal/core/domain/sessions"
	"github.com/go-hexagonal-practice/internal/core/ports"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type sessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) ports.SessionRepositoryPort {
	return &sessionRepository{db: db}
}

func (r *sessionRepository) Create(ctx context.Context, s *domain_sessions.UserSessions) error {
	// Map Domain -> GORM (Persistence)
	dbSession := &gorm_sessions.UserSessions{
		ID:               s.ID,
		UserID:           s.UserID,
		RefreshTokenHash: s.RefreshTokenHash,
		IPAddress:        s.IPAddress,
		UserAgent:        s.UserAgent,
		// Ensure other fields from your GORM model are mapped here
	}

	return r.db.WithContext(ctx).Create(dbSession).Error
}

func (r *sessionRepository) FindByID(ctx context.Context, sessionID uuid.UUID) (*domain_sessions.UserSessions, error) {
	var dbSession gorm_sessions.UserSessions

	if err := r.db.WithContext(ctx).First(&dbSession, "id = ?", sessionID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("session not found")
		}
		return nil, err
	}

	// Map GORM -> Domain
	return toDomainSession(&dbSession), nil
}

func (r *sessionRepository) Revoke(ctx context.Context, sessionID uuid.UUID) error {
	// Using your GORM model's 'IsRevoked' field
	result := r.db.WithContext(ctx).Model(&gorm_sessions.UserSessions{}).
		Where("id = ?", sessionID).
		Update("is_revoked", true)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("session not found")
	}

	return nil
}

// --- Mapper: GORM -> Domain ---

func toDomainSession(s *gorm_sessions.UserSessions) *domain_sessions.UserSessions {
	return &domain_sessions.UserSessions{
		ID:               s.ID,
		UserID:           s.UserID,
		RefreshTokenHash: s.RefreshTokenHash,
		IPAddress:        s.IPAddress,
		UserAgent:        s.UserAgent,
		// Add ExpiresAt, IsRevoked etc. if they exist in your domain struct
	}
}
