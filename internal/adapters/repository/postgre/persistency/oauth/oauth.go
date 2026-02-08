package oauth

import (
	"time"

	user_model "github.com/go-hexagonal-practice/internal/core/domain/user"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type OauthIdentities struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_user_provider;not null"`

	Provider   string `gorm:"type:varchar(50);uniqueIndex:idx_user_provider;uniqueIndex:idx_user_provider_id;not null"`
	ProviderID string `gorm:"type:varchar(255);uniqueIndex:idx_user_provider_id;not null"`

	ProfileData datatypes.JSON `gorm:"type:jsonb;not null"`

	CreatedAt time.Time `gorm:"type:timestamptz;default:now();not null"`
	UpdatedAt time.Time `gorm:"type:timestamptz;default:now();not null"`

	User *user_model.User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
