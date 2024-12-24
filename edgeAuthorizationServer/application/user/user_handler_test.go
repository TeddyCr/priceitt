package user

import (
	"context"
	"testing"

	"github.com/TeddyCr/priceitt/models/config"
	"github.com/TeddyCr/priceitt/models/generated/auth"
	"github.com/TeddyCr/priceitt/models/generated/createEntities"
	"github.com/TeddyCr/priceitt/models/generated/entities"
	goFernet "github.com/fernet/fernet-go"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/argon2"
	"github.com/TeddyCr/priceitt/edgeAuthorizationServer/infrastructure/fernet"
	repository "github.com/TeddyCr/priceitt/edgeAuthorizationServer/repository/database"
)

const (
	fernetString = "jwEMNW7F-XYPNe4s9jZRfv7Ra9rwMBgV-gDP4NxjAXA="
)

func TestUserHandler_Create(t *testing.T) {
	password := "passWord12345!!!"
	mockedRepository := repository.MockRepository{}
	salt := []byte("salt")
	fernetConfig := config.FernetConfig{
		Key:  fernetString,
		Salt: string(salt),
	}
	fernet.Initialize(fernetConfig)

	userHandler := NewUserHandler(mockedRepository)

	createUser := &createEntities.CreateUser{
		Name:            "test",
		Email:           "example@email.com",
		AuthType:        "basic",
		Password:        password,
		ConfirmPassword: password,
	}
	ctx := context.Background()
	user, err := userHandler.Create(ctx, createUser)
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}
	if user == nil {
		t.Fatalf("user is nil")
	}
	userEntity, _ := user.(*entities.User)
	auth, _ := userEntity.AuthenticationMechanism.(auth.Basic)
	assert.True(t, validatePassword(password, auth.Password))
	assert.NotNil(t, userEntity.ID)
}

func validatePassword(password string, encryptedPassword string) bool {
	fernetKey := goFernet.MustDecodeKeys(fernetString)
	decryptedPassword := goFernet.VerifyAndDecrypt([]byte(encryptedPassword), 0, fernetKey)
	hashedPassword := argon2.IDKey([]byte(password), []byte("salt"), 1, 64*1024, 4, 32)
	return string(hashedPassword) == string(decryptedPassword)
}
