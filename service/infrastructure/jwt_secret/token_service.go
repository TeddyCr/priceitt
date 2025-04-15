package jwt_secret

import (
	"context"

	repository "github.com/TeddyCr/priceitt/service/repository/database"
	auth_repository "github.com/TeddyCr/priceitt/service/repository/database/auth"
)

type ITokenService interface {
	IsTokenBlacklisted(ctx context.Context, token string, filter repository.QueryFilter) (bool, error)
}

var tokenService *TokenService

type TokenService struct {
	repository auth_repository.AuthRepository
}

func (t *TokenService) IsTokenBlacklisted(ctx context.Context, token string, filter repository.QueryFilter) (bool, error) {
	blacklistToken, err := t.repository.GetBlacklistToken(ctx, token, filter)
	if err != nil {
		return false, err
	}
	return blacklistToken != nil, nil
}

func GetTokenServiceInstance() *TokenService {
	if tokenService == nil {
		panic("TokenService not initialized. Call jwt_secret.InitializeTokenService() before using jwt_secret.GetTokenServiceInstance()")
	}

	return tokenService
}

func InitializeTokenService(authRepository auth_repository.AuthRepository) {
	lock.Lock()
	defer lock.Unlock()

	tokenService = &TokenService{
		repository: authRepository,
	}
}
