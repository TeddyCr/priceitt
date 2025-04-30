//go:build unit
// +build unit

package user

import (
	"context"
	"testing"
	"time"

	"github.com/TeddyCr/priceitt/service/handler"
	infrastructureFernet "github.com/TeddyCr/priceitt/service/infrastructure/fernet"
	"github.com/TeddyCr/priceitt/service/infrastructure/jwt_secret"
	"github.com/TeddyCr/priceitt/service/models/config"
	"github.com/TeddyCr/priceitt/service/models/generated"
	"github.com/TeddyCr/priceitt/service/models/generated/auth"
	"github.com/TeddyCr/priceitt/service/models/generated/createEntities"
	"github.com/TeddyCr/priceitt/service/models/generated/entities"
	"github.com/TeddyCr/priceitt/service/models/types"
	repository "github.com/TeddyCr/priceitt/service/repository/database"
	dbUtils "github.com/TeddyCr/priceitt/service/utils/database"
	"github.com/TeddyCr/priceitt/service/utils/fernet"
	"github.com/TeddyCr/priceitt/service/utils/jwt"
	goFernet "github.com/fernet/fernet-go"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/argon2"
)

// MockAuthRepository is a mock implementation of the auth repository
type MockAuthRepository struct {
}

func (m MockAuthRepository) Create(ctx context.Context, token generated.IEntity) error {
	return nil
}

func (m MockAuthRepository) CreateBlacklistToken(ctx context.Context, token generated.IEntity) error {
	return nil
}

func (m MockAuthRepository) Logout(ctx context.Context, token string, user entities.User) error {
	return nil
}

func (m MockAuthRepository) GetById(ctx context.Context, id string, filter repository.QueryFilter) (generated.IEntity, error) {
	return nil, nil
}

func (m MockAuthRepository) GetByName(ctx context.Context, name string, filter repository.QueryFilter) (generated.IEntity, error) {
	return nil, nil
}

func (m MockAuthRepository) UpdateById(ctx context.Context, id string, entity generated.IEntity, filter repository.QueryFilter) error {
	return nil
}

func (m MockAuthRepository) UpdateByName(ctx context.Context, name string, entity generated.IEntity, filter repository.QueryFilter) error {
	return nil
}

func (m MockAuthRepository) DeleteById(ctx context.Context, id string, filter repository.QueryFilter) error {
	return nil
}

func (m MockAuthRepository) DeleteByName(ctx context.Context, name string, filter repository.QueryFilter) error {
	return nil
}

func (m MockAuthRepository) List(ctx context.Context, filter repository.QueryFilter) ([]generated.IEntity, error) {
	return nil, nil
}

func (m MockAuthRepository) GetClient() dbUtils.Executor {
	return nil
}

// MockUserRepository is a mock implementation of the user repository
type MockUserRepository struct{}

func (m MockUserRepository) Create(ctx context.Context, createEntity generated.IEntity) error {
	return nil
}

func (m MockUserRepository) GetById(ctx context.Context, id string, filter repository.QueryFilter) (generated.IEntity, error) {
	return nil, nil
}

func (m MockUserRepository) GetByName(ctx context.Context, name string, filter repository.QueryFilter) (generated.IEntity, error) {
	password := "passWord12345!!!"
	fernetInstance := infrastructureFernet.GetInstance()
	hashedPassword := fernet.HashPasswordWithSalt(password, fernetInstance.Salt)
	encryptedPassword := fernet.EncryptAndSign(
		hashedPassword,
		fernetInstance.Key[0])
	userName := "test"
	userEmail := "test@example.com"
	user := &entities.User{
		ID:    uuid.New(),
		Name:  userName,
		Email: userEmail,
		AuthenticationMechanism: auth.Basic{
			Type:     "basic",
			Password: string(encryptedPassword),
		},
		CreatedAt: time.Now().UnixMilli(),
		UpdatedAt: time.Now().UnixMilli(),
	}
	return user, nil
}

func (m MockUserRepository) UpdateById(ctx context.Context, id string, entity generated.IEntity, filter repository.QueryFilter) error {
	return nil
}

func (m MockUserRepository) UpdateByName(ctx context.Context, name string, entity generated.IEntity, filter repository.QueryFilter) error {
	return nil
}

func (m MockUserRepository) DeleteById(ctx context.Context, id string, filter repository.QueryFilter) error {
	return nil
}

func (m MockUserRepository) DeleteByName(ctx context.Context, name string, filter repository.QueryFilter) error {
	return nil
}

func (m MockUserRepository) List(ctx context.Context, filter repository.QueryFilter) ([]generated.IEntity, error) {
	return nil, nil
}

func (m MockUserRepository) GetClient() dbUtils.Executor {
	return nil
}

const (
	fernetString = "jwEMNW7F-XYPNe4s9jZRfv7Ra9rwMBgV-gDP4NxjAXA="
)

func initFernet() {
	salt := []byte("salt")
	fernetConfig := config.FernetConfig{
		Key:  fernetString,
		Salt: string(salt),
	}
	infrastructureFernet.Initialize(fernetConfig)
}

func initJWT() {
	jwtConfig := config.SecretConfig{
		Secret:   "secret",
		Issuer:   "issuer",
		Audience: "audience",
	}
	jwt_secret.Initialize(jwtConfig)
}

func getUserHandler() handler.IHandler {
	mockedAuthRepository := MockAuthRepository{}
	mockedUserRepository := MockUserRepository{}

	return NewUserHandler(mockedUserRepository, mockedAuthRepository)
}

// Start of tests
func TestUserHandler_Create(t *testing.T) {
	password := "passWord12345!!!"
	initFernet()
	userHandler := getUserHandler()

	createUser := &createEntities.CreateUser{
		Name:     "test",
		Email:    "example@email.com",
		AuthType: "basic",
		AuthMechanism: auth.Basic{
			Type:            "basic",
			Password:        password,
			ConfirmPassword: password,
		},
	}
	ctx := context.Background()
	user, err := userHandler.Create(ctx, createUser, repository.QueryFilter{})
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

func TestUserHandler_Login(t *testing.T) {
	initFernet()
	initJWT()
	userHandler := getUserHandler()
	token, err := userHandler.(UserHandler).Login(context.Background(), auth.AuthEncapsulation{
		Type:     "basic",
		Username: "test",
		Data: json.RawMessage(auth.Basic{
			Password: "passWord12345!!!",
			Type:     "basic",
		}),
	})
	accessTokenEntity := token["access"].(*entities.JWToken)
	assert.NoError(t, err)
	assert.NotNil(t, accessTokenEntity.ID)
	assert.NotNil(t, accessTokenEntity.CreatedAt)
	assert.NotNil(t, accessTokenEntity.UpdatedAt)
	assert.Equal(t, accessTokenEntity.TokenType, types.AccessToken.String())
	assert.Equal(t, accessTokenEntity.Name, types.AccessToken.String())
	assert.True(t, accessTokenEntity.ExpirationDate > time.Now().UnixMilli())
	// since it was generated before we check the expiration time should be less than in 1 hour
	assert.True(t, accessTokenEntity.ExpirationDate <= time.Now().Add(time.Hour*1).UnixMilli())
}

func TestGetUser(t *testing.T) {
	password := "passWord12345!!!"
	createUser := &createEntities.CreateUser{
		Name:     "test",
		Email:    "example@email.com",
		AuthType: "basic",
		AuthMechanism: auth.Basic{
			Type:            "basic",
			Password:        password,
			ConfirmPassword: password,
		},
	}
	userHandler := getUserHandler()

	user := userHandler.(UserHandler).getUser(createUser, []byte(password))
	userEntity, _ := user.(*entities.User)
	assert.Equal(t, createUser.Name, userEntity.Name)
	assert.Equal(t, createUser.Email, userEntity.Email)
	assert.Equal(t, createUser.AuthType, userEntity.AuthenticationMechanism.(auth.Basic).Type)
	assert.Equal(t, createUser.Password, userEntity.AuthenticationMechanism.(auth.Basic).Password)
	assert.NotNil(t, userEntity.ID)
	assert.NotNil(t, userEntity.CreatedAt)
	assert.NotNil(t, userEntity.UpdatedAt)
}

func TestCreateTokens(t *testing.T) {
	initJWT()
	password := "passWord12345!!!"
	createUser := &createEntities.CreateUser{
		Name:     "test",
		Email:    "example@email.com",
		AuthType: "basic",
		AuthMechanism: auth.Basic{
			Type:            "basic",
			Password:        password,
			ConfirmPassword: password,
		},
	}
	userHandler := getUserHandler()

	user := userHandler.(UserHandler).getUser(createUser, []byte(password))
	userEntity, _ := user.(*entities.User)
	refreshToken := jwt.GetRefreshToken(userEntity.ID)
	assert.NotNil(t, refreshToken)
	assert.NotNil(t, refreshToken.ID)
	assert.NotNil(t, refreshToken.CreatedAt)
	assert.NotNil(t, refreshToken.UpdatedAt)
	assert.Equal(t, refreshToken.TokenType, types.RefreshToken.String())
	assert.Equal(t, refreshToken.Name, types.RefreshToken.String())
	assert.Equal(t, refreshToken.UserID, userEntity.ID)
	assert.NotNil(t, refreshToken.Token)
	assert.NotNil(t, refreshToken.ExpirationDate)

	accessToken := jwt.GetAccessToken(userEntity.ID)
	assert.NotNil(t, accessToken)
	assert.NotNil(t, accessToken.ID)
	assert.NotNil(t, accessToken.CreatedAt)
	assert.NotNil(t, accessToken.UpdatedAt)
	assert.Equal(t, accessToken.TokenType, types.AccessToken.String())
	assert.Equal(t, accessToken.Name, types.AccessToken.String())
	assert.Equal(t, accessToken.UserID, userEntity.ID)
	assert.NotNil(t, accessToken.Token)
	assert.NotNil(t, accessToken.ExpirationDate)
}

func validatePassword(password string, encryptedPassword string) bool {
	fernetKey := goFernet.MustDecodeKeys(fernetString)
	decryptedPassword := goFernet.VerifyAndDecrypt([]byte(encryptedPassword), 0, fernetKey)
	hashedPassword := argon2.IDKey([]byte(password), []byte("salt"), 1, 64*1024, 4, 32)
	return string(hashedPassword) == string(decryptedPassword)
}
