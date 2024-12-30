package repository

import (
	"context"
	"fmt"

	"github.com/TeddyCr/priceitt/edgeAuthorizationServer/infrastructure/database"
	repository "github.com/TeddyCr/priceitt/edgeAuthorizationServer/repository/database"
	"github.com/TeddyCr/priceitt/models/generated"
	utilDB "github.com/TeddyCr/priceitt/utils/database"
	"github.com/jmoiron/sqlx"
)

func NewUserRepository(dbContext database.IPersistenceDatabase) *UserRepository {
	return &UserRepository{dbContext: dbContext, client: dbContext.GetClient()}
}

// UserRepository is a struct that defines the methods that a user repository should implement
type UserRepository struct {
	dbContext database.IPersistenceDatabase
	client    *sqlx.DB
}

func (u *UserRepository) Create(ctx context.Context, user generated.IEntity) error {
	query := fmt.Sprintf(repository.InsertQuery, "users")
	err := utilDB.PerformEntityQuery(ctx, u.client, query, user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) GetById(ctx context.Context, id string) (generated.IEntity, error) { // (entities.User, error) {
	return nil, nil
}

func (u *UserRepository) GetByName(ctx context.Context, name string) (generated.IEntity, error) { // (entities.User, error) {
	return nil, nil
}

func (u *UserRepository) Update(ctx context.Context, entity generated.IEntity) error {
	return nil
}

func (u *UserRepository) Delete(ctx context.Context, id string) error {
	return nil
}

func (u *UserRepository) List(ctx context.Context) ([]generated.IEntity, error) { // ([]entities.User, error) {
	return nil, nil
}

func (u *UserRepository) GetClient() *sqlx.DB {
	return u.client
}
