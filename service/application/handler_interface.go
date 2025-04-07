package application

import (
	"context"

	"github.com/TeddyCr/priceitt/service/models/generated"
)

type IHandler interface {
	Create(ctx context.Context, create generated.ICreateEntity) (generated.IEntity, error)
}
