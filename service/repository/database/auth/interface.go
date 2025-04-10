package auth_repository

import (
	"context"

	"github.com/TeddyCr/priceitt/service/models/generated"
	"github.com/TeddyCr/priceitt/service/models/generated/entities"
	repository "github.com/TeddyCr/priceitt/service/repository/database"
	"github.com/TeddyCr/priceitt/service/utils/database"
)

type IAuthRepository interface {
	Create(ctx context.Context, createEntity generated.IEntity) error

	GetById(ctx context.Context, id string, filter repository.QueryFilter) (generated.IEntity, error)

	GetByName(ctx context.Context, name string, filter repository.QueryFilter) (generated.IEntity, error)

	UpdateById(ctx context.Context, id string, entity generated.IEntity, filter repository.QueryFilter) error

	UpdateByName(ctx context.Context, name string, entity generated.IEntity, filter repository.QueryFilter) error

	DeleteById(ctx context.Context, id string, filter repository.QueryFilter) error

	DeleteByName(ctx context.Context, name string, filter repository.QueryFilter) error

	List(ctx context.Context, filter repository.QueryFilter) ([]generated.IEntity, error)

	GetClient() database.Executor

	CreateBlacklistToken(ctx context.Context, token generated.IEntity) error

	Logout(ctx context.Context, token string, user entities.User) error
}
