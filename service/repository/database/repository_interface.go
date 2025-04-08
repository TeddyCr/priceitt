package repository

import (
	"context"

	"github.com/TeddyCr/priceitt/service/models/generated"
	"github.com/TeddyCr/priceitt/service/utils/database"
)

// IRepository is an interface that defines the methods that a repository should implement
type IDatabaseRepository interface {
	Create(ctx context.Context, createEntity generated.IEntity) error

	GetById(ctx context.Context, id string, filter QueryFilter) (generated.IEntity, error)

	GetByName(ctx context.Context, name string, filter QueryFilter) (generated.IEntity, error)

	UpdateById(ctx context.Context, id string, entity generated.IEntity, filter QueryFilter) error

	UpdateByName(ctx context.Context, name string, entity generated.IEntity, filter QueryFilter) error

	DeleteById(ctx context.Context, id string, filter QueryFilter) error

	DeleteByName(ctx context.Context, name string, filter QueryFilter) error

	List(ctx context.Context, filter QueryFilter) ([]generated.IEntity, error)

	GetClient() database.Executor
}
