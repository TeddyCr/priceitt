package jwt_secret

import (
	"strings"
	"sync"

	"github.com/TeddyCr/priceitt/models/config"
)

var lock = &sync.Mutex{}

type JWTSecret struct {
	Secret  []byte
	Issuer  string
	Audience []string
}

var instance *JWTSecret

func GetInstance() *JWTSecret {
	if instance == nil {
		panic("JWTSecret not initialized. Call jwt_secret.Initialize() before using jwt_secret.GetInstance()")
	}

	return instance
}

func Initialize(config config.SecretConfig) {
	lock.Lock()
	defer lock.Unlock()

	audience := strings.Split(config.Audience, ",")

	instance = &JWTSecret{
		Secret:   []byte(config.Secret),
		Issuer:   config.Issuer,
		Audience: audience,
	}
}
