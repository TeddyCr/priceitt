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

func TestValidateJsonSchemaInvalidLength(t *testing.T) {
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

func TestValidateJsonSchemaInvalidCharcters(t *testing.T) {
	userSchemaPath := "https://raw.githubusercontent.com/TeddyCr/priceitt/refs/heads/main/models/schema/createEntities/createUser.json";
	user := createEntities.CreateUser{
		Name: "Jane Smith",
		Email: "foo@bar.com",
		Password: "Password123456789",
		ConfirmPassword: "Password123456789",
	}

	validationResult, err := ValidateJsonSchema(userSchemaPath, user)
	
	assert.Nil(t, err)
	assert.False(t, validationResult.IsValid)
}

func TestBuildHttpModelPath(t *testing.T) {
	version := "v0.0.1-alpha"
	entityType := "createUser"
	extra := ""

	expected := "https://raw.githubusercontent.com/TeddyCr/priceitt/refs/tags/models/v0.0.1-alpha/schema/createEntities/createUser.json"
	buildHttpPathResult := buildHttpPath(version, entityType, extra)

	assert.Equal(t, expected, buildHttpPathResult)
}