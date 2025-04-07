package entities

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type JWToken struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	DisplayName    string    `json:"displayName"`
	UpdatedAt      int64     `json:"updatedAt"`
	CreatedAt      int64     `json:"createdAt"`
	Token          string    `json:"token"`
	TokenType      string    `json:"tokenType"`
	ExpirationDate int64     `json:"expirationDate"`
	UserID         uuid.UUID `json:"userId"`
	DeviceID       uuid.UUID `json:"deviceId"`
	IP             string    `json:"ip"`
}

func (jwt JWToken) GetID() uuid.UUID {
	return jwt.ID
}

func (jwt JWToken) GetName() string {
	return jwt.Name
}

func (jwt JWToken) GetDisplayName() string {
	return jwt.DisplayName
}

func (jwt JWToken) GetDescription() string {
	return jwt.Description
}

func (jwt JWToken) GetUpdatedAt() int64 {
	return jwt.UpdatedAt
}

func (jwt JWToken) GetCreatedAt() int64 {
	return jwt.CreatedAt
}

func (jwt JWToken) ToJson() ([]byte, error) {
	return json.Marshal(jwt)
}

func (jwt *JWToken) Bind(r *http.Request) error {
	return nil
}

func (jwt *JWToken) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
