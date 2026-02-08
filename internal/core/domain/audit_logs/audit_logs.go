package auditlogs

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type AuditLogs struct {
	ID        uint64
	CreatedAt time.Time

	UserID     *uuid.UUID
	Action     string     // LOGIN_SUCCESS, LOGIN_FAILED, PASSWORD_CHANGED, COUNTRY_UPDATED_ML and etc
	EntityType string     // USER, SEESION, BILLING and etc
	EntityID   *uuid.UUID // The ID of the object being changed

	IPAddress *string
	UserAgent *string
	Metadata  datatypes.JSON
}
