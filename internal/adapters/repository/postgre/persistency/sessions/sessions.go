package sessions

import (
	"encoding/json"
	"time"

	"github.com/go-hexagonal-practice/internal/adapters/repository/postgre/persistency/user"
	domain_sessions "github.com/go-hexagonal-practice/internal/core/domain/sessions"
	"github.com/go-hexagonal-practice/internal/core/ports"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type UserSessions struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID uuid.UUID `gorm:"type:uuid;not null;index:idx_sessions_user_active,where:is_revoked = false"`

	RefreshTokenHash string `gorm:"type:varchar(255);not null;index:idx_sessions_token"`

	IPAddress   string          `gorm:"type:inet;not null"`
	UserAgent   *string         `gorm:"type:text"`
	Device      *string         `gorm:"type:varchar(255)"`
	GeoLocation *datatypes.JSON `gorm:"type:jsonb"`

	IsRevoked    bool      `gorm:"type:boolean;default:false;not null;index:idx_sessions_user_active,where:is_revoked = false"`
	ExpiresAt    time.Time `gorm:"type:timestamptz;index;not null"`
	CreatedAt    time.Time `gorm:"type:timestamptz;default:now();not null"`
	LastActiveAt time.Time `gorm:"type:timestamptz;default:now();not null"`

	User *user.User `gorm:"foreignKey:UserID;reference:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (u *UserSessions) ToDomain(logger ports.Logger) *domain_sessions.UserSessions {
	domainSession := domain_sessions.UserSessions{
		ID:               u.ID,
		UserID:           u.UserID,
		RefreshTokenHash: u.RefreshTokenHash,
		IPAddress:        u.IPAddress,
		UserAgent:        u.UserAgent,
		Device:           u.Device,
		// GeoLocation:      u.GeoLocation,
		ExpiresAt: u.ExpiresAt,
	}
	if u.GeoLocation != nil {
		var geo domain_sessions.GeoLocation
		if err := json.Unmarshal(*u.GeoLocation, &geo); err == nil {
			domainSession.GeoLocation = &geo
		} else {
			logger.Error("Geo Location Mapping Error left the column value empty", "error", err.Error())
		}
	}
	return &domainSession
}
