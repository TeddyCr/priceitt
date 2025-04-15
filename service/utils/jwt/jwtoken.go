package jwt

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/TeddyCr/priceitt/service/infrastructure/jwt_secret"
	"github.com/TeddyCr/priceitt/service/models/generated/entities"
	"github.com/TeddyCr/priceitt/service/models/types"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func CreateJWT(expiration int, sub string, tokenType string) (string, error) {
	jwtInstance := jwt_secret.GetInstance()
	key := jwtInstance.Secret
	issuer := jwtInstance.Issuer
	claims := jwt.MapClaims{
		"iss":  issuer,
		"exp":  time.Now().Add(time.Hour * time.Duration(expiration)).UnixMilli(),
		"iat":  time.Now().UnixMilli(),
		"sub":  sub,
		"type": tokenType,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func DecodeAndValidateJWT(encodedToken string) (jwt.MapClaims, error) {
	decodedToken, err := DecodeJWT(encodedToken)
	if err != nil {
		return nil, err
	}
	jwtInstance := jwt_secret.GetInstance()
	key := jwtInstance.Secret
	token, err := jwt.Parse(decodedToken, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	expired, err := IsExpired(claims)
	if err != nil {
		return nil, err
	}
	if expired {
		return nil, errors.New("token expired")
	}
	if !IsAccessToken(claims) {
		return nil, errors.New("token is not a valid access token")
	}
	return claims, nil
}

func DecodeJWT(token string) (string, error) {
	decodedToken, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return "", err
	}
	return string(decodedToken), nil
}

func IsExpired(claims jwt.MapClaims) (bool, error) {
	exp, err := claims.GetExpirationTime()
	if err != nil {
		return false, err
	}

	return exp.Before(time.Now()), nil
}

func IsAccessToken(claims jwt.MapClaims) bool {
	tokenType, ok := claims["type"].(string)
	if !ok {
		return false
	}
	return tokenType == types.TokenType(types.AccessToken).String()
}

func GetRefreshToken(userId uuid.UUID) *entities.JWToken {
	var expiration = 999999
	expirationEnv := os.Getenv("REFRESH_EXPIRATION")
	if expirationEnv != "" {
		expiration, _ = strconv.Atoi(expirationEnv)
	}
	refreshToken, err := CreateJWT(expiration, userId.String(), "refresh")
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
		UserID:         userId,
		// TODO: get device id and ip from request
		DeviceID: uuid.New(),
		IP:       "",
	}
}

func GetAccessToken(userId uuid.UUID) *entities.JWToken {
	var expiration = 1
	expirationEnv := os.Getenv("ACCESS_EXPIRATION")
	if expirationEnv != "" {
		expiration, _ = strconv.Atoi(expirationEnv)
	}
	accessToken, err := CreateJWT(expiration, userId.String(), "access")
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
		UserID:         userId,
		// TODO: get device id and ip from request
		DeviceID: uuid.New(),
		IP:       "",
	}
}
