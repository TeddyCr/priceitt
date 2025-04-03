package generated

import (
	"net/http"

	"github.com/google/uuid"
)

type IEntity interface {
	GetID() uuid.UUID
	GetName() string
	GetDisplayName() string
	GetDescription() string
	GetUpdatedAt() int64
	GetCreatedAt() int64
	ToJson() ([]byte, error)
	Bind(r *http.Request) error
	Render(w http.ResponseWriter, r *http.Request) error
}
