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

func NewAuthRepository(dbContext database.IPersistenceDatabase) *AuthRepository {
	return &AuthRepository{dbContext: dbContext, client: dbContext.GetClient()}
}

// AuthRepository is a struct that defines the methods that an auth repository should implement
type AuthRepository struct {
	dbContext database.IPersistenceDatabase
	client    *sqlx.DB
}

func (a *AuthRepository) Create(ctx context.Context, user generated.IEntity) error {
	query := fmt.Sprintf(repository.InsertQuery, "tokens")
	err := utilDB.PerformEntityQuery(ctx, a.client, query, user)
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthRepository) GetById(ctx context.Context, id string) (generated.IEntity, error) { // (entities.User, error) {
	return nil, nil
}

func (a *AuthRepository) GetByName(ctx context.Context, name string) (generated.IEntity, error) { // (entities.User, error) {
	return nil, nil
}

func (a *AuthRepository) Update(ctx context.Context, entity generated.IEntity) error {
	return nil
}

func (a *AuthRepository) Delete(ctx context.Context, id string) error {
	return nil
}

func (a *AuthRepository) List(ctx context.Context) ([]generated.IEntity, error) { // ([]entities.User, error) {
	return nil, nil
}

func (a *AuthRepository) GetClient() *sqlx.DB {
	return a.client
}
