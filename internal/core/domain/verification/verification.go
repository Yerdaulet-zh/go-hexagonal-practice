package verification

import (
	"time"

	user_model "github.com/go-hexagonal-practice/internal/core/domain/user"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/datatypes"
)

type Verifications struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Type      string
	TokenHash string

	Metadata *datatypes.JSON // Store context ("triggered by an ml risk engine")

	IsUsed    bool
	ExpiresAt time.Time
	CreatedAt time.Time

	User *user_model.User
}

type UsersMFASecrets struct {
	UserID      uuid.UUID
	SecretKey   string
	BackupCodes pq.StringArray
	CreatedAt   time.Time

	User *user_model.User
}
