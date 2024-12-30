package entities

import (
	"encoding/json"
	"net/http"

	"github.com/TeddyCr/priceitt/models/generated/auth"
	"github.com/google/uuid"
)

type User struct {
	ID                      uuid.UUID   `json:"id"`
	Name                    string      `json:"name"`
	Description             string      `json:"description"`
	DisplayName             string      `json:"displayName"`
	UpdatedAt               int64       `json:"updatedAt"`
	CreatedAt               int64       `json:"createdAt"`
	Email                   string      `json:"email"`
	Image50                 string      `json:"image50"`
	AuthenticationMechanism interface{} `json:"authenticationMechanism"` // map to oneOf
}

func (u User) GetID() uuid.UUID {
	return u.ID
}

func (u User) GetName() string {
	return u.Name
}

func (u User) GetDisplayName() string {
	return u.DisplayName
}

func (u User) GetDescription() string {
	return u.Description
}

func (u User) GetUpdatedAt() int64 {
	return u.UpdatedAt
}

func (u User) GetCreatedAt() int64 {
	return u.CreatedAt
}

func (u User) ToJson() ([]byte, error) {
	return json.Marshal(u)
}

func (u *User) Bind(r *http.Request) error {
	return nil
}

func (u *User) Render(w http.ResponseWriter, r *http.Request) error {
	authMechanism := u.AuthenticationMechanism.(auth.BaseAuthMechanism)
	if authMechanism.GetAuthType() == "basic" {
		basicAuth := authMechanism.(auth.Basic)
		basicAuth.Password = ""
		u.AuthenticationMechanism = basicAuth
	}
	return nil
}
