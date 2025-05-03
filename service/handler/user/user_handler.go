package user

import (
	"context"
	"crypto/subtle"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/TeddyCr/priceitt/service/handler"
	"github.com/TeddyCr/priceitt/service/models/generated"
	"github.com/TeddyCr/priceitt/service/models/generated/auth"
	"github.com/TeddyCr/priceitt/service/models/generated/createEntities"
	"github.com/TeddyCr/priceitt/service/models/generated/entities"
	"github.com/TeddyCr/priceitt/service/models/types"
	"github.com/TeddyCr/priceitt/service/serializer"
	goFernet "github.com/fernet/fernet-go"
	"github.com/go-chi/httplog/v2"
	"github.com/google/uuid"

	infrastructureFernet "github.com/TeddyCr/priceitt/service/infrastructure/fernet"
	repository "github.com/TeddyCr/priceitt/service/repository/database"
	auth_repository "github.com/TeddyCr/priceitt/service/repository/database/auth"
	user_repository "github.com/TeddyCr/priceitt/service/repository/database/user"
	"github.com/TeddyCr/priceitt/service/utils/fernet"
	"github.com/TeddyCr/priceitt/service/utils/jwt"
	"google.golang.org/api/idtoken"
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

	var encryptedSecret []byte
	switch createUser.AuthType {
	case "basic":
		authMechanism, err := serializer.DeSerBasicAuthMechanism(createUser.AuthMechanism.(map[string]interface{}))
		if err != nil {
			return nil, err
		}
		err = createUser.ValidatePassword(authMechanism)
		if err != nil {
			return nil, err
		}
		hashedSecret := fernet.HashPasswordWithSalt(authMechanism.Password, c.fernetInstance.Salt)
		encryptedSecret = fernet.EncryptAndSign(hashedSecret, c.fernetInstance.Key[0])
	case "google":
		authMechanism, err := serializer.DeSerGoogleAuthMechanism(createUser.AuthMechanism.(map[string]interface{}))
		if err != nil {
			return nil, err
		}
		err = c.handleGoogleAuth(authMechanism)
		if err != nil {
			return nil, err
		}
		hashedSecret := fernet.HashPasswordWithSalt(authMechanism.IdToken, c.fernetInstance.Salt)
		encryptedSecret = fernet.EncryptAndSign(hashedSecret, c.fernetInstance.Key[0])
	}

	user := c.getUser(createUser, encryptedSecret)
	err := c.DatabaseRepository.Create(ctx, user)
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

func (c UserHandler) Login(ctx context.Context, authEncapsuler auth.AuthEncapsulation) (map[string]generated.IEntity, error) {
	var logger = httplog.LogEntry(ctx)
	var user generated.IEntity
	var err error
	authMechanism, err := c.getAuthMechanism(authEncapsuler)
	if err != nil {
		return nil, err
	}
	user, err = c.DatabaseRepository.GetByName(ctx, authEncapsuler.GetUsername(), *repository.NewQueryFilter(nil))
	if err != nil {
		user, err = c.DatabaseRepository.(*user_repository.UserRepository).GetByEmail(ctx, authEncapsuler.GetUsername(), *repository.NewQueryFilter(nil))
		if err != nil && errors.Is(err, sql.ErrNoRows) {
			logger.Debug("failed to get user", "error", err)
			return nil, fmt.Errorf("user [%s] not found", authEncapsuler.GetUsername())
		} else if err != nil {
			logger.Debug("failed to get user", "error", err)
			return nil, err
		}
	}
	userEntity, ok := user.(*entities.User)
	if !ok {
		return nil, errors.New("failed to cast to entities.User")
	}
	switch authEncapsuler.GetAuthType() {
	case "basic":
		err = c.handleBasicAuth(userEntity, authMechanism.(auth.Basic))
	case "google":
		err = c.handleGoogleAuth(authMechanism.(auth.Google))
	default:
		return nil, errors.New("invalid auth type")
	}
	if err != nil {
		return nil, err
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
	var authMechanism auth.BaseAuthMechanism
	switch createUser.AuthType {
	case "basic":
		authMechanism = auth.Basic{
			Type:     "basic",
			Password: string(encryptedPassword),
		}
	case "google":
		authMechanism = auth.Google{
			Type:     "google",
			IdToken:  createUser.AuthMechanism.(map[string]interface{})["idToken"].(string),
			Audience: os.Getenv("GOOGLE_CLIENT_ID"),
		}
	}
	now := time.Now().UnixMilli()
	return &entities.User{
		ID:                      uuid.New(),
		Name:                    createUser.Name,
		CreatedAt:               now,
		UpdatedAt:               now,
		Email:                   createUser.Email,
		AuthenticationMechanism: authMechanism,
	}
}

func (c UserHandler) validatePassword(encryptedPassword string, password string) bool {
	decryptedPassword := goFernet.VerifyAndDecrypt([]byte(encryptedPassword), 0, c.fernetInstance.Key)
	hashedPassword := fernet.HashPasswordWithSalt(password, c.fernetInstance.Salt)
	return subtle.ConstantTimeCompare(hashedPassword, decryptedPassword) == 1
}

func (c UserHandler) getAuthMechanism(authEncapsuler auth.AuthEncapsulation) (auth.BaseAuthMechanism, error) {
	switch authEncapsuler.GetAuthType() {
	case "basic":
		var basicAuth auth.Basic
		err := json.Unmarshal(authEncapsuler.GetData(), &basicAuth)
		if err != nil {
			return nil, err
		}
		return basicAuth, nil
	case "google":
		var googleAuth auth.Google
		err := json.Unmarshal(authEncapsuler.GetData(), &googleAuth)
		if err != nil {
			return nil, err
		}
		return googleAuth, nil
	default:
		return nil, errors.New("invalid auth type")
	}
}

func (c UserHandler) handleBasicAuth(userEntity *entities.User, basicAuth auth.Basic) error {
	authJson, err := json.Marshal(userEntity.AuthenticationMechanism)
	if err != nil {
		return err
	}
	var auth auth.Basic
	if err := json.Unmarshal(authJson, &auth); err != nil {
		return err
	}
	if !c.validatePassword(auth.Password, basicAuth.GetPassword()) {
		return errors.New("invalid password")
	}
	return nil
}

func (c UserHandler) handleGoogleAuth(authMechanism auth.Google) error {
	validator, err := idtoken.NewValidator(context.Background())
	if err != nil {
		return err
	}
	_, err = validator.Validate(context.Background(), authMechanism.GetIdToken(), authMechanism.GetAudience())
	if err != nil {
		return err
	}
	return nil
}
