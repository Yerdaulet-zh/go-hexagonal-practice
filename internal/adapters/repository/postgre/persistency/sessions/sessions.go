package sessions

import (
	"time"

	user_model "github.com/go-hexagonal-practice/internal/core/domain/user"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

/*
CREATE TABLE user_sessions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,

    -- Security Tokens
    refresh_token_hash VARCHAR(255) NOT NULL, -- Don't store raw tokens!

    -- Device/Context Info (Critical for Analysis)
    ip_address INET NOT NULL, -- Postgres optimized IP storage
    user_agent TEXT,
    device_fingerprint VARCHAR(255), -- Client-side generated unique ID
    geo_location JSONB, -- Lat/Long snapshot at login

    -- Lifecycle
    is_revoked BOOLEAN DEFAULT FALSE,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    last_active_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Index for fast lookups during auth middleware checks
CREATE INDEX idx_sessions_user_active ON user_sessions(user_id) WHERE is_revoked = FALSE;
CREATE INDEX idx_sessions_token ON user_sessions(refresh_token_hash);
*/

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

	User *user_model.User `gorm:"foreignKey:UserID;reference:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
