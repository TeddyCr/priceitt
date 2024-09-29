package json

import (
	"testing"
	"github.com/TeddyCr/priceitt/models/generated/createEntities"
	"github.com/stretchr/testify/assert"
)

func TestValidateJsonSchemaValid(t *testing.T) {
	userSchemaPath := "https://raw.githubusercontent.com/TeddyCr/priceitt/refs/heads/main/models/schema/createEntities/createUser.json";
	user := createEntities.CreateUser{
		Name: "Jane Smith",
		Email: "foo@bar.com",
		Password: "Password123!@#!!",
		ConfirmPassword: "Password123!@#!!",
	}

	validationResult, err := ValidateJsonSchema(userSchemaPath, user)
	
	assert.Nil(t, err)
	assert.True(t, validationResult.IsValid)
}

func TestValidateJsonSchemaInvalid(t *testing.T) {
	userSchemaPath := "https://raw.githubusercontent.com/TeddyCr/priceitt/refs/heads/main/models/schema/createEntities/createUser.json";
	user := createEntities.CreateUser{
		Name: "Jane Smith",
		Email: "foo@bar.com",
		Password: "Password12",
		ConfirmPassword: "Password12",
	}

	validationResult, err := ValidateJsonSchema(userSchemaPath, user)
	
	assert.Nil(t, err)
	assert.False(t, validationResult.IsValid)
}