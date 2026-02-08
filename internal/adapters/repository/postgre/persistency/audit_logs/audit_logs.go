package auditlogs

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type AuditLogs struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement:false"`
	CreatedAt time.Time `gorm:"primaryKey;type:timestamptz;default:now()"`

	UserID     *uuid.UUID `gorm:"type:uuid;index"`
	Action     string     `gorm:"type:varchar(50);not null"` // LOGIN_SUCCESS, LOGIN_FAILED, PASSWORD_CHANGED, COUNTRY_UPDATED_ML and etc
	EntityType string     `gorm:"type:varchar(50);not null"` // USER, SEESION, BILLING and etc
	EntityID   *uuid.UUID `gorm:"type:uuid"`                 // The ID of the object being changed

	IPAddress *string        `gorm:"type:inet"`
	UserAgent *string        `gorm:"type:text"`
	Metadata  datatypes.JSON `gorm:"type:jsonb;default:{};not null"`
}
