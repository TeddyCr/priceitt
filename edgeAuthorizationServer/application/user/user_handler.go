package user

import (
	"context"
	"log"
	"time"

	"github.com/TeddyCr/priceitt/models/generated"
	"github.com/TeddyCr/priceitt/models/generated/auth"
	"github.com/TeddyCr/priceitt/models/generated/createEntities"
	"github.com/TeddyCr/priceitt/models/generated/entities"
	"github.com/fernet/fernet-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/argon2"

	dbRepo "priceitt.xyz/edgeAuthorizationServer/repository/database"
)

type UserHandler struct {
	DatabaseRepository dbRepo.IDatabaseRepository
	FernetKey []*fernet.Key
	Salt []byte
}

func (c *UserHandler) Create(ctx context.Context, createUser createEntities.CreateUser) (generated.IEntity, error) {
	err := createUser.ValidatePassword()
	if err != nil {
		log.Fatalf("failed to validate password: %v", err)
		return nil, err
	}
	hashedPassword := argon2.IDKey(
		[]byte(createUser.Password),
		c.Salt,
		1,
		64 * 1024,
		4,
		32)
	token, err := fernet.EncryptAndSign(hashedPassword, c.FernetKey[0])
	if err != nil {
		log.Fatalf("failed to encrypt password: %v", err)
		return nil, err
	}
	user := c.getUser(createUser, token)
	c.DatabaseRepository.Create(ctx, user)
	return user, nil
}

func (c UserHandler) getUser(createUser createEntities.CreateUser, encryptedPassword []byte) entities.User {
	now := time.Now()
	return entities.User{
		BaseEntity: entities.BaseEntity{
			ID: uuid.New(),
			Name: createUser.Name,
			CreatedAt: now,
			UpdatedAt: now,
		},
		Email: createUser.Email,
		AuthenticationMechanism: auth.Basic{
			Type: "basic",
			Password: string(encryptedPassword),
		},
	}
}
