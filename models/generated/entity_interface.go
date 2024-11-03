package generated

import (
	"time"

	"github.com/google/uuid"
)

type IEntity interface {
	GetID() uuid.UUID
	GetName() string
	GetDisplayName() string
	GetDescription() string
	GetUpdatedAt() time.Time
	GetCreatedAt() time.Time
	ToJson() ([]byte, error)
	JsonToString() (string, error)
}
