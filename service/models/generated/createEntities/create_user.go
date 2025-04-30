package createEntities

import (
	"errors"
	"net/http"

	"github.com/TeddyCr/priceitt/service/models/generated/auth"
)

type CreateUser struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	AuthType        string `json:"authType"`
	AuthMechanism   any 	`json:"authMechanism"`
}

func (c *CreateUser) ValidatePassword() error {
	if err := c.ValidatePasswordLength(); err != nil {
		return err
	}
	if err := c.ValidateConfirmPassword(); err != nil {
		return err
	}
	if err := c.ValidatePasswordCharacters(); err != nil {
		return err
	}
	return nil
}

func (c *CreateUser) ValidatePasswordLength() error {
	if len(c.AuthMechanism.(auth.Basic).Password) < 16 {
		return errors.New("password must be at least 16 characters")
	}
	return nil
}

func (c *CreateUser) ValidateConfirmPassword() error {
	if c.AuthMechanism.(auth.Basic).Password != c.AuthMechanism.(auth.Basic).ConfirmPassword {
		return errors.New("passwords do not match")
	}
	return nil
}

func (c *CreateUser) ValidatePasswordCharacters() error {
	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSpecial := false

	for _, char := range c.AuthMechanism.(auth.Basic).Password {
		switch {
		case 'A' <= char && char <= 'Z':
			hasUpper = true
		case 'a' <= char && char <= 'z':
			hasLower = true
		case '0' <= char && char <= '9':
			hasNumber = true
		case char == '!' || char == '@' || char == '#' || char == '$' || char == '%' || char == '^' || char == '&' || char == '*':
			hasSpecial = true
		}
	}

	if !hasUpper || !hasLower || !hasNumber || !hasSpecial {
		return errors.New("password must contain at least one uppercase letter, one lowercase letter, one number, and one special character (!@#$%^&*)")
	}

	return nil
}

func (c CreateUser) GetName() string {
	return c.Name
}

func (c *CreateUser) Bind(r *http.Request) error {
	return nil
}

func (c *CreateUser) Render(w http.ResponseWriter, r *http.Request) error {
	switch c.AuthType {
		case "basic":
			c.AuthMechanism.(*auth.Basic).Password = ""
			c.AuthMechanism.(*auth.Basic).ConfirmPassword = ""
		case "google":
			c.AuthMechanism.(*auth.Google).IdToken = ""
	}
	return nil
}
