package verification

import (
	"time"

	user_model "github.com/go-hexagonal-practice/internal/core/domain/user"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/datatypes"
)

type Verifications struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID    uuid.UUID `gorm:"type:uuid;index;not null;"`
	Type      string    `gorm:"type:verification_type;not null"`
	TokenHash string    `gorm:"type:varchar(255);not null"`

	Metadata *datatypes.JSON `gorm:"type:jsonb"` // Store context ("triggered by an ml risk engine")

	IsUsed    bool      `gorm:"type:boolean;default:false;not null"`
	ExpiresAt time.Time `gorm:"type:timestamptz;not null"`
	CreatedAt time.Time `gorm:"type:timestamptz;default:now();not null"`

	User *user_model.User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type UsersMFASecrets struct {
	UserID      uuid.UUID      `gorm:"type:uuid;primaryKey"`
	SecretKey   string         `gorm:"type:varchar(255);not null;"`
	BackupCodes pq.StringArray `gorm:"type:text[];not null"`
	CreatedAt   time.Time      `gorm:"type:timestamptz;default:now();not null"`

	User *user_model.User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
