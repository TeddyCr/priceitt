package user

import (
	"context"
	"fmt"
	"time"

	"github.com/TeddyCr/priceitt/models/generated"
	"github.com/TeddyCr/priceitt/models/generated/auth"
	"github.com/TeddyCr/priceitt/models/generated/createEntities"
	"github.com/TeddyCr/priceitt/models/generated/entities"
	goFernet "github.com/fernet/fernet-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/argon2"

	"github.com/TeddyCr/priceitt/edgeAuthorizationServer/application"
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
	createUser, ok := createEntity.(*createEntities.CreateUser)
	if !ok {
		panic("failed to cast to createEntities.CreateUser")
	}
	err := createUser.ValidatePassword()
	if err != nil {
		panic(fmt.Sprintf("failed to validate password: %v", err))
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
		panic(fmt.Sprintf("failed to encrypt password: %v", err))
	}
	user := c.getUser(createUser, token)
	err = c.DatabaseRepository.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c UserHandler) getUser(createUser *createEntities.CreateUser, encryptedPassword []byte) generated.IEntity {
	now := time.Now().UnixMilli()
	return &entities.User{
		ID:        uuid.New(),
		Name:      createUser.Name,
		CreatedAt: now,
		UpdatedAt: now,
		Email: createUser.Email,
		AuthenticationMechanism: auth.Basic{
			Type:     "basic",
			Password: string(encryptedPassword),
		},
	}
}
