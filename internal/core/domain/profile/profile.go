package profile

import (
	"time"

	user_model "github.com/go-hexagonal-practice/internal/core/domain/user"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type UserProfile struct {
	UserID        uuid.UUID
	FirstName     string
	LastName      *string
	CountryCode   *string
	CountrySource *string
	AvatarURL     *string
	Preferences   *datatypes.JSON
	UpdatedAt     time.Time

	User *user_model.User
}

type UserProfileHistory struct {
	HistoryId     uuid.UUID
	UserID        uuid.UUID
	FirstName     string
	LastName      *string
	CountryCode   *string
	CountrySource *string
	AvatarURL     *string
	ChangedAt     time.Time
	Operation     string

	User *user_model.User
}
