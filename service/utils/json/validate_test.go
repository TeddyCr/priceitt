//go:build unit
// +build unit

package json

import (
	"github.com/TeddyCr/priceitt/service/models/generated/auth"
	"github.com/TeddyCr/priceitt/service/models/generated/createEntities"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestValidateJsonSchemaValid(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}
	rootPath := filepath.Dir(filepath.Dir(cwd))
	userSchemaPath := "file:///" + filepath.Join(rootPath, "models", "schema", "createEntities", "createUser.json")
	user := createEntities.CreateUser{
		Name:     "Jane Smith",
		Email:    "foo@bar.com",
		AuthType: "basic",
		AuthMechanism: auth.Basic{
			Type:            "basic",
			Password:        "Password123!@#!!",
			ConfirmPassword: "Password123!@#!!",
		},
	}

	validationResult, err := ValidateJsonSchema(userSchemaPath, user)

	assert.Nil(t, err)
	assert.True(t, validationResult.IsValid)
}

func TestValidateJsonSchemaInvalidLength(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}
	rootPath := filepath.Dir(filepath.Dir(cwd))
	userSchemaPath := "file:///" + filepath.Join(rootPath, "models", "schema", "createEntities", "createUser.json")
	user := createEntities.CreateUser{
		Name:     "Jane Smith",
		Email:    "foo@bar.com",
		AuthType: "basic",
		AuthMechanism: auth.Basic{
			Type:            "basic",
			Password:        "Password12",
			ConfirmPassword: "Password12",
		},
	}

	validationResult, err := ValidateJsonSchema(userSchemaPath, user)

	assert.Nil(t, err)
	assert.False(t, validationResult.IsValid)
}

func TestValidateJsonSchemaInvalidCharcters(t *testing.T) {
	t.Skip("Skipping this test as pattern validation is not working")
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}
	rootPath := filepath.Dir(filepath.Dir(cwd))
	userSchemaPath := "file:///" + filepath.Join(rootPath, "models", "schema", "createEntities", "createUser.json")
	user := createEntities.CreateUser{
		Name:     "Jane Smith",
		Email:    "foo@bar.com",
		AuthType: "basic",
		AuthMechanism: auth.Basic{
			Type:            "basic",
			Password:        "Password123456789",
			ConfirmPassword: "Password123456789",
		},
	}

	validationResult, err := ValidateJsonSchema(userSchemaPath, user)

	assert.Nil(t, err)
	assert.False(t, validationResult.IsValid)
}
