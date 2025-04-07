package repository

import (
	"context"

	"github.com/TeddyCr/priceitt/service/models/generated"
	"github.com/TeddyCr/priceitt/service/utils/database"
)

type MockRepository struct {
}

func (m MockRepository) Create(ctx context.Context, createEntity generated.IEntity) error {
	return nil
}

func (m MockRepository) GetById(ctx context.Context, id string, filter QueryFilter) (generated.IEntity, error) {
	return nil, nil
}

func (m MockRepository) GetByName(ctx context.Context, name string, filter QueryFilter) (generated.IEntity, error) {
	return nil, nil
}

func (m MockRepository) UpdateById(ctx context.Context, id string, entity generated.IEntity, filter QueryFilter) error {
	return nil
}

func (m MockRepository) UpdateByName(ctx context.Context, name string, entity generated.IEntity, filter QueryFilter) error {
	return nil
}

func (m MockRepository) DeleteById(ctx context.Context, id string, filter QueryFilter) error {
	return nil
}

func (m MockRepository) DeleteByName(ctx context.Context, name string, filter QueryFilter) error {
	return nil
}

func (m MockRepository) List(ctx context.Context, filter QueryFilter) ([]generated.IEntity, error) {
	return nil, nil
}

func (m MockRepository) GetClient() database.Executor {
	return nil
}
