package user_repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/TeddyCr/priceitt/service/infrastructure/database"
	repository "github.com/TeddyCr/priceitt/service/repository/database"
	"github.com/TeddyCr/priceitt/service/models/generated"
	"github.com/TeddyCr/priceitt/service/models/generated/entities"
	utilDB "github.com/TeddyCr/priceitt/service/utils/database"
	"github.com/jmoiron/sqlx"
)

type fn func (ctx context.Context, db *sqlx.DB, query string, name string) (*sql.Row, error)


func NewUserRepository(dbContext database.IPersistenceDatabase) *UserRepository {
	return &UserRepository{dbContext: dbContext, client: dbContext.GetClient()}
}

// UserRepository is a struct that defines the methods that a user repository should implement
type UserRepository struct {
	dbContext database.IPersistenceDatabase
	client    *sqlx.DB
}

func (u UserRepository) Create(ctx context.Context, user generated.IEntity) error {
	query := fmt.Sprintf(repository.InsertQuery, "users")
	err := utilDB.PerformEntityQuery(ctx, u.client, query, user)
	if err != nil {
		return err
	}
	return nil
}

func (u UserRepository) GetById(ctx context.Context, id string) (generated.IEntity, error) { // (entities.User, error) {
	return nil, nil
}

func (u UserRepository) GetByName(ctx context.Context, name string) (generated.IEntity, error) { // (entities.User, error) {
	query := fmt.Sprintf(repository.GetByName, "users")
	user, err := u.getByNameOrEmail(ctx, utilDB.PerformSelectScalarQuery, query, name)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserRepository) GetByEmail(ctx context.Context, email string) (generated.IEntity, error) { // (entities.User, error) {
	query := fmt.Sprintf(GetByEmail, "users")
	user, err := u.getByNameOrEmail(ctx, utilDB.PerformSelectScalarQuery, query, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserRepository) Update(ctx context.Context, entity generated.IEntity) error {
	return nil
}

func (u UserRepository) Delete(ctx context.Context, id string) error {
	return nil
}

func (u UserRepository) List(ctx context.Context) ([]generated.IEntity, error) { // ([]entities.User, error) {
	return nil, nil
}

func (u UserRepository) GetClient() *sqlx.DB {
	return u.client
}

func (u UserRepository) getByNameOrEmail(ctx context.Context, f fn, query string, name string) (generated.IEntity, error) {
	row, err := f(ctx, u.client, query, name)
	if err != nil {
		return nil, err
	}
	var jsonData []byte
	err = row.Scan(&jsonData)
	if err != nil {
		return nil, err
	}
	var user entities.User
	err = json.Unmarshal(jsonData, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
