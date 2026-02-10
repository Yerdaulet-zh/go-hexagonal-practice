package sessions

import (
	"time"

	user_model "github.com/go-hexagonal-practice/internal/core/domain/user"
	"github.com/google/uuid"
)

type GeoLocation struct {
	City        string  `json:"city"`
	Country     string  `json:"country"`
	CountryCode string  `json:"country_code"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	TimeZone    string  `json:"timezone"`
}

type UserSessions struct {
	ID     uuid.UUID
	UserID uuid.UUID

	RefreshTokenHash string

	IPAddress   string
	UserAgent   *string
	Device      *string
	GeoLocation *GeoLocation // Datatypes JSON

	IsRevoked    bool
	ExpiresAt    time.Time
	CreatedAt    time.Time
	LastActiveAt time.Time

	// Maybe in future it may be deleted.
	User *user_model.User
}
