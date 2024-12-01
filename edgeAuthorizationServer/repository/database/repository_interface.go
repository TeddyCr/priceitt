package repository

import (
	"context"

	"github.com/TeddyCr/priceitt/models/generated"
	"github.com/jmoiron/sqlx"
)

// IRepository is an interface that defines the methods that a repository should implement
type IDatabaseRepository interface {
	Create(ctx context.Context, createEntity generated.IEntity) error

	GetById(ctx context.Context, id string) (generated.IEntity, error)

	GetByName(ctx context.Context, name string) (generated.IEntity, error)

	Update(ctx context.Context, entity generated.IEntity) error

	Delete(ctx context.Context, id string) error

	List(ctx context.Context) ([]generated.IEntity, error)

	getClient() *sqlx.DB
}
