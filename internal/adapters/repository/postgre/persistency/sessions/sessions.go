package sessions

import (
	"time"

	"github.com/go-hexagonal-practice/internal/adapters/repository/postgre/persistency/user"
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
