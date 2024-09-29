package entities

import (
	"time"
)

type User struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Image50 string `json:"image50"`
	Description string `json:"description"`
	AuthenticationMechanism interface{} `json:"authenticationMechanism"` // map to oneOf
}

