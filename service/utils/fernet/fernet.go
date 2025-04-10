package fernet

import (
	"fmt"

	goFernet "github.com/fernet/fernet-go"
	"golang.org/x/crypto/argon2"
)

func HashPasswordWithSalt(password string, salt []byte) []byte {
	return argon2.IDKey(
		[]byte(password),
		salt,
		1,
		64*1024,
		4,
		32)
}

func HashPassword(password string) []byte {
	return HashPasswordWithSalt(password, []byte(""))
}

func EncryptAndSign(password []byte, key *goFernet.Key) []byte {
	token, err := goFernet.EncryptAndSign(password, key)
	if err != nil {
		panic(fmt.Sprintf("failed to encrypt password: %v", err))
	}
	return token
}
