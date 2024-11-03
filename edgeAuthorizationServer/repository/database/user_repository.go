package database

import (
	"context"
	"fmt"

	"github.com/TeddyCr/priceitt/models/generated/entities"
	utilDB "github.com/TeddyCr/priceitt/utils/database"
	"github.com/jmoiron/sqlx"
	"priceitt.xyz/edgeAuthorizationServer/infrastructure/database"
)

func NewUserRepository(dbContext database.IPersistenceDatabase) *UserRepository {
	return &UserRepository{dbContext: dbContext, client: dbContext.GetClient()}
}

// UserRepository is a struct that defines the methods that a user repository should implement
type UserRepository struct {
	dbContext database.IPersistenceDatabase
	client *sqlx.DB
}

func (u *UserRepository) Create(ctx context.Context, user entities.User) error {
	query := fmt.Sprintf(InsertQuery, "users")
	err := utilDB.PerformEntityQuery(ctx, u.client, query, user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) GetById(ctx context.Context, id string) error { // (entities.User, error) {
	return nil
}

func (u *UserRepository) GetByName(ctx context.Context, name string) error { // (entities.User, error) {
	return nil
}

func (u *UserRepository) Update(ctx context.Context, entity entities.User) error {
	return nil
}

func (u *UserRepository) Delete(ctx context.Context, id string) error {
	return nil
}

func (u *UserRepository) List(ctx context.Context) error { // ([]entities.User, error) {
	return nil
}
