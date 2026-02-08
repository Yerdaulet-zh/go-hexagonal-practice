package oauth

import (
	"time"

	user_model "github.com/go-hexagonal-practice/internal/core/domain/user"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type OauthIdentities struct {
	ID     uuid.UUID
	UserID uuid.UUID

	Provider   string
	ProviderID string

	ProfileData datatypes.JSON

	CreatedAt time.Time
	UpdatedAt time.Time

	User *user_model.User
}
