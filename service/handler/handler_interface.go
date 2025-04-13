package handler

import (
	"context"

	"github.com/TeddyCr/priceitt/service/models/generated"
	repository "github.com/TeddyCr/priceitt/service/repository/database"
)

type IHandler interface {
	Create(ctx context.Context, create generated.ICreateEntity, filter repository.QueryFilter) (generated.IEntity, error)
	GetByName(ctx context.Context, name string, filter repository.QueryFilter) (generated.IEntity, error)
}
