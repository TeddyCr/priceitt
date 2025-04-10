//go:build unit
// +build unit

package fernet

import (
	"testing"

	goFernet "github.com/fernet/fernet-go"
	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	password := "password"
	hashedPassword := HashPassword(password)
	assert.NotNil(t, hashedPassword)
	assert.NotEqual(t, password, hashedPassword)
	assert.Equal(t, hashedPassword, HashPassword(password))

	// Test with salt
	salt := []byte("salt")
	hashedPasswordWithSalt := HashPasswordWithSalt(password, salt)
	assert.NotNil(t, hashedPasswordWithSalt)
	assert.NotEqual(t, hashedPassword, hashedPasswordWithSalt)
	assert.Equal(t, hashedPasswordWithSalt, HashPasswordWithSalt(password, salt))
}

func TestEncryptAndSign(t *testing.T) {
	password := "password"
	key, err := goFernet.DecodeKey("jwEMNW7F-XYPNe4s9jZRfv7Ra9rwMBgV-gDP4NxjAXA=")
	assert.NoError(t, err)
	encryptedPassword := EncryptAndSign([]byte(password), key)
	assert.NotNil(t, encryptedPassword)
	assert.NotEqual(t, password, encryptedPassword)
}
