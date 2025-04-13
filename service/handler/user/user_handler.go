package user

import (
	"context"
	"crypto/subtle"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/TeddyCr/priceitt/service/handler"
	"github.com/TeddyCr/priceitt/service/models/generated"
	"github.com/TeddyCr/priceitt/service/models/generated/auth"
	"github.com/TeddyCr/priceitt/service/models/generated/createEntities"
	"github.com/TeddyCr/priceitt/service/models/generated/entities"
	"github.com/TeddyCr/priceitt/service/models/types"
	goFernet "github.com/fernet/fernet-go"
	"github.com/go-chi/httplog/v2"
	"github.com/google/uuid"

	infrastructureFernet "github.com/TeddyCr/priceitt/service/infrastructure/fernet"
	repository "github.com/TeddyCr/priceitt/service/repository/database"
	auth_repository "github.com/TeddyCr/priceitt/service/repository/database/auth"
	user_repository "github.com/TeddyCr/priceitt/service/repository/database/user"
	"github.com/TeddyCr/priceitt/service/utils/fernet"
	"github.com/TeddyCr/priceitt/service/utils/jwt"
)

func NewUserHandler(databaseRepository repository.IDatabaseRepository, authRepository auth_repository.IAuthRepository) handler.IHandler {
	return UserHandler{
		DatabaseRepository: databaseRepository,
		AuthRepository:     authRepository,
		fernetInstance:     infrastructureFernet.GetInstance(),
	}
}

type UserHandler struct {
	DatabaseRepository repository.IDatabaseRepository
	AuthRepository     auth_repository.IAuthRepository
	fernetInstance     *infrastructureFernet.Fernet
}

func (c UserHandler) Create(ctx context.Context, createEntity generated.ICreateEntity, filter repository.QueryFilter) (generated.IEntity, error) {
	createUser, ok := createEntity.(*createEntities.CreateUser)
	if !ok {
		panic("failed to cast to createEntities.CreateUser")
	}
	err := createUser.ValidatePassword()
	if err != nil {
		panic(fmt.Sprintf("failed to validate password: %v", err))
	}
	hashedPassword := fernet.HashPasswordWithSalt(createUser.Password, c.fernetInstance.Salt)
	token := fernet.EncryptAndSign(hashedPassword, c.fernetInstance.Key[0])
	user := c.getUser(createUser, token)
	err = c.DatabaseRepository.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c UserHandler) GetByName(ctx context.Context, name string, filter repository.QueryFilter) (generated.IEntity, error) {
	user, err := c.DatabaseRepository.GetByName(ctx, name, filter)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c UserHandler) Login(ctx context.Context, basicAuth auth.BasicAuth) (map[string]generated.IEntity, error) {
	var logger = httplog.LogEntry(ctx)
	var user generated.IEntity
	var err error
	user, err = c.DatabaseRepository.GetByName(ctx, basicAuth.Username, *repository.NewQueryFilter(nil))
	if err != nil {
		user, err = c.DatabaseRepository.(*user_repository.UserRepository).GetByEmail(ctx, basicAuth.Username, *repository.NewQueryFilter(nil))
		if err != nil && errors.Is(err, sql.ErrNoRows) {
			logger.Debug("failed to get user", "error", err)
			return nil, fmt.Errorf("user [%s] not found", basicAuth.Username)
		} else if err != nil {
			logger.Debug("failed to get user", "error", err)
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
	refresh := jwt.GetRefreshToken(userEntity.ID)
	access := jwt.GetAccessToken(userEntity.ID)
	err = c.AuthRepository.Create(ctx, refresh)
	if err != nil {
		return nil, err
	}
	return map[string]generated.IEntity{
		"access":  access,
		"refresh": refresh,
	}, nil
}

func (c UserHandler) Logout(ctx context.Context) (string, error) {
	var logger = httplog.LogEntry(ctx)
	jwtContextValues := ctx.Value("jwtContextValues").(types.JWTContextValues)
	token, ok := jwtContextValues.Get("token").(string)
	if !ok {
		logger.Debug(fmt.Sprintf("failed to get token from context: %v", jwtContextValues))
		return "", errors.New("failed to get token from context")
	}
	userId, ok := jwtContextValues.Get("userId").(string)
	if !ok {
		logger.Debug(fmt.Sprintf("failed to get user id from context: %v", jwtContextValues))
		return "", errors.New("failed to get user id from context")
	}

	user, err := c.DatabaseRepository.GetById(ctx, userId, *repository.NewQueryFilter(nil))
	if err != nil {
		logger.Debug(fmt.Sprintf("failed to get user: %v", err))
		return "", err
	}
	err = c.AuthRepository.Logout(ctx, token, *user.(*entities.User))
	if err != nil {
		logger.Debug(fmt.Sprintf("failed to logout: %v", err))
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

func (c UserHandler) validatePassword(encryptedPassword string, password string) bool {
	decryptedPassword := goFernet.VerifyAndDecrypt([]byte(encryptedPassword), 0, c.fernetInstance.Key)
	hashedPassword := fernet.HashPasswordWithSalt(password, c.fernetInstance.Salt)
	return subtle.ConstantTimeCompare(hashedPassword, decryptedPassword) == 1
}
