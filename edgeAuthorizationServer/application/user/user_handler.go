package user

import (
	"context"
	"log"
	"time"

	"github.com/TeddyCr/priceitt/models/generated"
	"github.com/TeddyCr/priceitt/models/generated/auth"
	"github.com/TeddyCr/priceitt/models/generated/createEntities"
	"github.com/TeddyCr/priceitt/models/generated/entities"
	goFernet "github.com/fernet/fernet-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/argon2"

	"priceitt.github.com/TeddyCr/priceitt/edgeAuthorizationServer/application"
	"github.com/TeddyCr/priceitt/edgeAuthorizationServer/infrastructure/fernet"
	dbRepo "github.com/TeddyCr/priceitt/edgeAuthorizationServer/repository/database"
)

func NewUserHandler(databaseRepository dbRepo.IDatabaseRepository) application.IHandler {
	return UserHandler{
		DatabaseRepository: databaseRepository,
		fernetInstance:     fernet.GetInstance(),
	}
}

type UserHandler struct {
	DatabaseRepository dbRepo.IDatabaseRepository
	fernetInstance     *fernet.Fernet
}

func (c UserHandler) Create(ctx context.Context, createEntity generated.ICreateEntity) (generated.IEntity, error) {
	createUser, ok := createEntity.(createEntities.CreateUser)
	if !ok {
		log.Fatalf("failed to cast to createEntities.CreateUser")
	}
	err := createUser.ValidatePassword()
	if err != nil {
		log.Fatalf("failed to validate password: %v", err)
		return nil, err
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
		log.Fatalf("failed to encrypt password: %v", err)
		return nil, err
	}
	user := c.getUser(createUser, token)
	c.DatabaseRepository.Create(ctx, user)
	return user, nil
}

func (c UserHandler) getUser(createEntity generated.ICreateEntity, encryptedPassword []byte) generated.IEntity {
	createUser, ok := createEntity.(createEntities.CreateUser)
	if !ok {
		log.Fatalf("failed to cast to createEntities.CreateUser")
	}
	now := time.Now()
	return entities.User{
		BaseEntity: entities.BaseEntity{
			ID:        uuid.New(),
			Name:      createUser.Name,
			CreatedAt: now,
			UpdatedAt: now,
		},
		Email: createUser.Email,
		AuthenticationMechanism: auth.Basic{
			Type:     "basic",
			Password: string(encryptedPassword),
		},
	}
}
