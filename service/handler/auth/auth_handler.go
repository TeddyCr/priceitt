package auth

import (
	"context"
	"errors"

	"github.com/TeddyCr/priceitt/service/handler"
	"github.com/TeddyCr/priceitt/service/models/generated"
	"github.com/TeddyCr/priceitt/service/models/generated/entities"
	"github.com/TeddyCr/priceitt/service/models/types"
	"github.com/TeddyCr/priceitt/service/utils/jwt"
	"github.com/google/uuid"

	infrastructureFernet "github.com/TeddyCr/priceitt/service/infrastructure/fernet"
	repository "github.com/TeddyCr/priceitt/service/repository/database"
	authRepository "github.com/TeddyCr/priceitt/service/repository/database/auth"
)

func NewAuthHandler(databaseRepository repository.IDatabaseRepository) handler.IHandler {
	return AuthHandler{
		DatabaseRepository: databaseRepository,
		fernetInstance:     infrastructureFernet.GetInstance(),
	}
}

type AuthHandler struct {
	DatabaseRepository repository.IDatabaseRepository
	fernetInstance     *infrastructureFernet.Fernet
}

func (a AuthHandler) Create(ctx context.Context, createEntity generated.ICreateEntity, filter repository.QueryFilter) (generated.IEntity, error) {
	return nil, nil
}

func (a AuthHandler) GetByName(ctx context.Context, name string, filter repository.QueryFilter) (generated.IEntity, error) {
	token, err := a.DatabaseRepository.GetByName(ctx, name, filter)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (a AuthHandler) CreateAccessToken(ctx context.Context, userIdStr string) (generated.IEntity, error) {
	jwtContextValues := ctx.Value("jwtContextValues").(types.JWTContextValues)
	token, ok := jwtContextValues.Get("token").(string)
	if !ok {
		return nil, errors.New("failed to get token from context")
	}
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		return nil, err
	}
	accessToken := jwt.GetAccessToken(userId)
	err = a.DatabaseRepository.(*authRepository.AuthRepository).CreateBlacklistToken(ctx, &entities.JWToken{
		ID:        uuid.New(),
		IP:        "",
		Token:     token,
		TokenType: types.TokenType(types.AccessToken).String(),
		UserID:    userId,
		Name:      "access",
	})
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

func (a AuthHandler) GetBlacklistToken(ctx context.Context, token string, filter repository.QueryFilter) (generated.IEntity, error) {
	blacklistToken, err := a.DatabaseRepository.(*authRepository.AuthRepository).GetBlacklistToken(ctx, token, filter)
	if err != nil {
		return nil, err
	}
	return blacklistToken, nil
}
