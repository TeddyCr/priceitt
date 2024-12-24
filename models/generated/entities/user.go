package entities

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	DisplayName string `json:"displayName"`
	UpdatedAt  	int64 `json:"updatedAt"`
	CreatedAt 	int64 `json:"createdAt"`
	Email string `json:"email"`
	Image50 string `json:"image50"`
	AuthenticationMechanism interface{} `json:"authenticationMechanism"` // map to oneOf
}

func (b User) GetID() uuid.UUID {
	return b.ID
}

func (b User) GetName() string {
	return b.Name
}

func (b User) GetDisplayName() string {
	return b.DisplayName
}

func (b User) GetDescription() string {
	return b.Description
}

func (b User) GetUpdatedAt() int64 {
	return b.UpdatedAt
}

func (b User) GetCreatedAt() int64 {
	return b.CreatedAt
}

func (b User) ToJson() ([]byte, error) {
	return json.Marshal(b)
}

func (b *User) Bind(r *http.Request) error {
	return nil
}

func (b *User) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}