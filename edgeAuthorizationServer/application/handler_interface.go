package application

import "github.com/TeddyCr/priceitt/models/generated"

type IHandler interface {
	Create(create generated.ICreateEntity) (generated.IEntity, error)
}
