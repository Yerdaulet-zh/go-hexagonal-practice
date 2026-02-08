package sessions

import (
	"time"

	user_model "github.com/go-hexagonal-practice/internal/core/domain/user"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type UserSessions struct {
	ID     uuid.UUID
	UserID uuid.UUID

	RefreshTokenHash string

	IPAddress   string
	UserAgent   *string
	Device      *string
	GeoLocation *datatypes.JSON

	IsRevoked    bool
	ExpiresAt    time.Time
	CreatedAt    time.Time
	LastActiveAt time.Time

	User *user_model.User
}
