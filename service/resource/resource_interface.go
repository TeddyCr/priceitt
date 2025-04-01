package resource

import (
	"github.com/go-chi/chi/v5"
)

// IUserResource is the interface for the user resource.
type IResource interface {
	Routes() chi.Router
}
