package repository

import (
	"context"

	"github.com/TeddyCr/priceitt/service/models/generated"
	"github.com/jmoiron/sqlx"
)

type MockRepository struct {
}

func (m MockRepository) Create(ctx context.Context, createEntity generated.IEntity) error {
	return nil
}

func (m MockRepository) GetById(ctx context.Context, id string) (generated.IEntity, error) {
	return nil, nil
}

func (m MockRepository) GetByName(ctx context.Context, name string) (generated.IEntity, error) {
	return nil, nil
}

func (m MockRepository) Update(ctx context.Context, entity generated.IEntity) error {
	return nil
}

func (m MockRepository) Delete(ctx context.Context, id string) error {
	return nil
}

func (m MockRepository) List(ctx context.Context) ([]generated.IEntity, error) {
	return nil, nil
}

func (m MockRepository) GetClient() *sqlx.DB {
	return nil
}
