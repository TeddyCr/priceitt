package user

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/TeddyCr/priceitt/service/models/generated"
	"github.com/TeddyCr/priceitt/service/models/generated/auth"
	"github.com/TeddyCr/priceitt/service/models/generated/createEntities"
	"github.com/TeddyCr/priceitt/service/models/generated/entities"
	"github.com/TeddyCr/priceitt/service/models/types"
	goFernet "github.com/fernet/fernet-go"
	"github.com/go-chi/httplog/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/argon2"

	"github.com/TeddyCr/priceitt/service/application"
	"github.com/TeddyCr/priceitt/service/infrastructure/fernet"
	dbRepo "github.com/TeddyCr/priceitt/service/repository/database"
	auth_repository "github.com/TeddyCr/priceitt/service/repository/database/auth"
	user_repository "github.com/TeddyCr/priceitt/service/repository/database/user"
)

func NewUserHandler(databaseRepository dbRepo.IDatabaseRepository, authRepository *auth_repository.AuthRepository) application.IHandler {
	return UserHandler{
		DatabaseRepository: databaseRepository,
		AuthRepository:     authRepository,
		fernetInstance:     fernet.GetInstance(),
	}
}

type UserHandler struct {
	DatabaseRepository dbRepo.IDatabaseRepository
	AuthRepository     *auth_repository.AuthRepository
	fernetInstance     *fernet.Fernet
}

func (c UserHandler) Create(ctx context.Context, createEntity generated.ICreateEntity) (generated.IEntity, error) {
	createUser, ok := createEntity.(*createEntities.CreateUser)
	if !ok {
		panic("failed to cast to createEntities.CreateUser")
	}
	err := createUser.ValidatePassword()
	if err != nil {
		panic(fmt.Sprintf("failed to validate password: %v", err))
	}
	hashedPassword := argon2.IDKey(
		[]byte(createUser.Password),
		c.fernetInstance.Salt,
		1,
		64*1024,
		4,
		32)
	token, err := goFernet.EncryptAndSign(hashedPassword, c.fernetInstance.Key[0])
	if err != nil {
		panic(fmt.Sprintf("failed to encrypt password: %v", err))
	}
	user := c.getUser(createUser, token)
	err = c.DatabaseRepository.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c UserHandler) Login(ctx context.Context, basicAuth auth.BasicAuth) (generated.IEntity, error) {
	var logger = httplog.LogEntry(ctx)
	var user generated.IEntity
	var err error
	user, err = c.DatabaseRepository.GetByName(ctx, basicAuth.Username, *dbRepo.NewQueryFilter(nil))
	if err != nil {
		user, err = c.DatabaseRepository.(*user_repository.UserRepository).GetByEmail(ctx, basicAuth.Username, *dbRepo.NewQueryFilter(nil))
		if err != nil && errors.Is(err, sql.ErrNoRows) {
			logger.Error("failed to get user", "error", err)
			return nil, fmt.Errorf("user [%s] not found", basicAuth.Username)
		} else if err != nil {
			return nil, err
		}
	}
	userEntity, ok := user.(*entities.User)
	if !ok {
		return nil, errors.New("failed to cast to entities.User")
	}
	authJson, err := json.Marshal(userEntity.AuthenticationMechanism)
	if err != nil {
		return nil, err
	}
	var auth auth.Basic
	if err := json.Unmarshal(authJson, &auth); err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("failed to cast to auth.Basic")
	}
	if !c.validatePassword(auth.Password, basicAuth.Password) {
		return nil, errors.New("invalid password")
	}
	refresh := c.createRefreshToken(userEntity)
	access := c.createAccessToken(userEntity)
	err = c.AuthRepository.Create(ctx, refresh)
	if err != nil {
		return nil, err
	}
	return access, nil
}

func (c UserHandler) Logout(ctx context.Context) (string, error) {
	var logger = httplog.LogEntry(ctx)
	jwtContextValues := ctx.Value("jwtContextValues").(types.JWTContextValues)
	token, ok := jwtContextValues.Get("token").(string)
	if !ok {
		logger.Error(fmt.Sprintf("failed to get token from context: %v", jwtContextValues))
		return "", errors.New("failed to get token from context")
	}
	userId, ok := jwtContextValues.Get("userId").(string)
	if !ok {
		logger.Error(fmt.Sprintf("failed to get user id from context: %v", jwtContextValues))
		return "", errors.New("failed to get user id from context")
	}

	user, err := c.DatabaseRepository.GetById(ctx, userId, *dbRepo.NewQueryFilter(nil))
	if err != nil {
		logger.Error(fmt.Sprintf("failed to get user: %v", err))
		return "", err
	}
	err = c.AuthRepository.Logout(ctx, token, *user.(*entities.User))
	if err != nil {
		logger.Error(fmt.Sprintf("failed to logout: %v", err))
		return "", err
	}

	return "Logout successful", nil
}

func (c UserHandler) getUser(createUser *createEntities.CreateUser, encryptedPassword []byte) generated.IEntity {
	now := time.Now().UnixMilli()
	return &entities.User{
		ID:        uuid.New(),
		Name:      createUser.Name,
		CreatedAt: now,
		UpdatedAt: now,
		Email:     createUser.Email,
		AuthenticationMechanism: auth.Basic{
			Type:     "basic",
			Password: string(encryptedPassword),
		},
	}
}

func (c UserHandler) createRefreshToken(userEntity *entities.User) *entities.JWToken {
	var expiration = 999999
	expirationEnv := os.Getenv("REFRESH_EXPIRATION")
	if expirationEnv != "" {
		expiration, _ = strconv.Atoi(expirationEnv)
	}
	refreshToken, err := CreateJWT(expiration, userEntity.ID.String(), "refresh")
	if err != nil {
		panic(fmt.Sprintf("failed to create refresh token: %v", err))
	}

	return &entities.JWToken{
		ID:             uuid.New(),
		Name:           types.TokenType(types.RefreshToken).String(),
		CreatedAt:      time.Now().UnixMilli(),
		UpdatedAt:      time.Now().UnixMilli(),
		TokenType:      types.TokenType(types.RefreshToken).String(),
		Token:          refreshToken,
		ExpirationDate: time.Now().Add(time.Hour * time.Duration(expiration)).UnixMilli(),
		UserID:         userEntity.ID,
		// TODO: get device id and ip from request
		DeviceID: uuid.New(),
		IP:       "",
	}
}

func (c UserHandler) createAccessToken(userEntity *entities.User) *entities.JWToken {
	var expiration = 1
	expirationEnv := os.Getenv("ACCESS_EXPIRATION")
	if expirationEnv != "" {
		expiration, _ = strconv.Atoi(expirationEnv)
	}
	accessToken, err := CreateJWT(expiration, userEntity.ID.String(), "access")
	if err != nil {
		panic(fmt.Sprintf("failed to create access token: %v", err))
	}

	return &entities.JWToken{
		ID:             uuid.New(),
		CreatedAt:      time.Now().UnixMilli(),
		UpdatedAt:      time.Now().UnixMilli(),
		TokenType:      types.TokenType(types.AccessToken).String(),
		Name:           types.TokenType(types.AccessToken).String(),
		Token:          accessToken,
		ExpirationDate: time.Now().Add(time.Hour * time.Duration(expiration)).UnixMilli(),
		UserID:         userEntity.ID,
		// TODO: get device id and ip from request
		DeviceID: uuid.New(),
		IP:       "",
	}
}

func (c UserHandler) validatePassword(encryptedPassword string, password string) bool {
	decryptedPassword := goFernet.VerifyAndDecrypt([]byte(encryptedPassword), 0, c.fernetInstance.Key)
	hashedPassword := argon2.IDKey([]byte(password), []byte(c.fernetInstance.Salt), 1, 64*1024, 4, 32)
	return string(hashedPassword) == string(decryptedPassword)
}
