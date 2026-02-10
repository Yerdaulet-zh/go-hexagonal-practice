package postgre

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	repo_sessions "github.com/go-hexagonal-practice/internal/adapters/repository/postgre/persistency/sessions"
	repo_user "github.com/go-hexagonal-practice/internal/adapters/repository/postgre/persistency/user"
	domain_sessions "github.com/go-hexagonal-practice/internal/core/domain/sessions"
	domain_user "github.com/go-hexagonal-practice/internal/core/domain/user"
	"github.com/go-hexagonal-practice/internal/core/ports"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type UserRepo struct {
	db     *gorm.DB
	logger ports.Logger
}

func NewUserRepository(db *gorm.DB, logger ports.Logger) *UserRepo {
	return &UserRepo{
		db:     db,
		logger: logger,
	}
}

func (repo *UserRepo) GetUserByEmail(ctx context.Context, email string) (*domain_user.User, error) {
	userRecord, err := gorm.G[repo_user.User](repo.db).Where("email = ?", email).Take(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &domain_user.User{}, fmt.Errorf("Such user by %s not found", email)
		}
		return &domain_user.User{}, err // Likelly Database Error
	}
	return userRecord.ToDomain(), nil

}

func (repo *UserRepo) CreateUser(ctx context.Context, user *domain_user.User, userCredentials *domain_user.UserCredentials, userSessions *domain_sessions.UserSessions) (*domain_sessions.UserSessions, error) {
	u := repo_user.User{
		Email:      user.Email,
		UserStatus: user.UserStatus,
		// IsMFAEnabled: user.IsMFAEnabled,
	}
	var sessionRecord repo_sessions.UserSessions
	err := repo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// User Creation
		if err := gorm.G[repo_user.User](tx).Create(ctx, &u); err != nil {
			return err
		}

		// User Credentials
		creds := repo_user.UserCredentials{
			UserID:       u.ID,
			PasswordHash: userCredentials.PasswordHash,
		}
		if err := gorm.G[repo_user.UserCredentials](tx).Create(ctx, &creds); err != nil {
			return err
		}

		// User Session
		sessionRecord = repo_sessions.UserSessions{
			UserID:           u.ID,
			RefreshTokenHash: userSessions.RefreshTokenHash,
			IPAddress:        userSessions.IPAddress,
			UserAgent:        userSessions.UserAgent,
			Device:           userSessions.Device,
			ExpiresAt:        userSessions.ExpiresAt,
		}

		// Convert Domain GeoLocation Struct -> JSONB for Database
		if userSessions.GeoLocation != nil {
			geoBytes, err := json.Marshal(userSessions.GeoLocation)
			if err != nil {
				return fmt.Errorf("failed to marshal geolocation: %w", err)
			} else {
				geoJSON := datatypes.JSON(geoBytes)
				sessionRecord.GeoLocation = &geoJSON
			}
		}
		if err := gorm.G[repo_sessions.UserSessions](tx).Create(ctx, &sessionRecord); err != nil {
			return err
		}

		// User Profile

		return nil
	})
	if err != nil {
		return nil, err
	}

	return sessionRecord.ToDomain(repo.logger), nil
}
