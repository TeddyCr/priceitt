package jwt

import (
	"errors"
	"time"

	"github.com/TeddyCr/priceitt/service/infrastructure/jwt_secret"
	"github.com/golang-jwt/jwt/v5"
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

func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	jwtInstance := jwt_secret.GetInstance()
	key := jwtInstance.Secret
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
