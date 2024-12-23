package createEntities

import "errors"

type CreateUser struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	AuthType string `json:"authType"`
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
	if len(c.Password) < 16 {
		return errors.New("password must be at least 16 characters")
	}
	return nil
}

func (c *CreateUser) ValidateConfirmPassword() error {
	if c.Password != c.ConfirmPassword {
		return errors.New("passwords do not match")
	}
	return nil
}

func (c *CreateUser) ValidatePasswordCharacters() error {
	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSpecial := false

	for _, char := range c.Password {
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

