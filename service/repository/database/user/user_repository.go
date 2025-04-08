package user_repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/TeddyCr/priceitt/service/infrastructure/database"
	"github.com/TeddyCr/priceitt/service/models/generated"
	"github.com/TeddyCr/priceitt/service/models/generated/entities"
	repository "github.com/TeddyCr/priceitt/service/repository/database"
	utilDB "github.com/TeddyCr/priceitt/service/utils/database"
)

type fn func(ctx context.Context, db utilDB.Executor, query string, args ...any) (*sql.Row, error)

func NewUserRepository(dbContext database.IPersistenceDatabase) *UserRepository {
	return &UserRepository{dbContext: dbContext, client: dbContext.GetClient()}
}

// UserRepository is a struct that defines the methods that a user repository should implement
type UserRepository struct {
	dbContext database.IPersistenceDatabase
	client    utilDB.Executor
}

func (u UserRepository) Create(ctx context.Context, user generated.IEntity) error {
	query := fmt.Sprintf(repository.InsertQuery, "users")
	err := utilDB.PerformEntityQuery(ctx, u.client, query, user)
	if err != nil {
		return err
	}
	return nil
}

func (u UserRepository) GetById(ctx context.Context, id string, filter repository.QueryFilter) (generated.IEntity, error) {
	query := fmt.Sprintf(repository.GetById, "users", filter.String())
	filter.Filter["id"] = id
	row, err := utilDB.PerformSelectScalarQuery(ctx, u.client, query, filter.Args()...)
	if err != nil {
		return nil, err
	}
	return marshalRow(row)
}

func (u UserRepository) GetByName(ctx context.Context, name string, filter repository.QueryFilter) (generated.IEntity, error) {
	query := fmt.Sprintf(repository.GetByName, "users", filter.String())
	filter.Filter["name"] = name
	user, err := u.getByNameOrEmail(ctx, utilDB.PerformSelectScalarQuery, query, filter.Args()...)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserRepository) GetByEmail(ctx context.Context, email string, filter repository.QueryFilter) (generated.IEntity, error) {
	query := fmt.Sprintf(GetByEmail, "users", filter.String())
	filter.Filter["email"] = email
	user, err := u.getByNameOrEmail(ctx, utilDB.PerformSelectScalarQuery, query, filter.Args()...)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserRepository) UpdateById(ctx context.Context, id string, entity generated.IEntity, filter repository.QueryFilter) error {
	return nil
}

func (u UserRepository) UpdateByName(ctx context.Context, name string, entity generated.IEntity, filter repository.QueryFilter) error {
	return nil
}

func (u UserRepository) DeleteById(ctx context.Context, id string, filter repository.QueryFilter) error {
	query := fmt.Sprintf(repository.DeleteById, "users", filter.String())
	filter.Filter["id"] = id
	_, err := u.client.ExecContext(ctx, query, filter.Args()...)
	if err != nil {
		return err
	}
	return nil
}

func (u UserRepository) DeleteByName(ctx context.Context, name string, filter repository.QueryFilter) error {
	query := fmt.Sprintf(repository.DeleteByName, "users", filter.String())
	filter.Filter["name"] = name
	_, err := u.client.ExecContext(ctx, query, filter.Args()...)
	if err != nil {
		return err
	}
	return nil
}

func (u UserRepository) List(ctx context.Context, filter repository.QueryFilter) ([]generated.IEntity, error) {
	return nil, nil
}

func (u UserRepository) GetClient() utilDB.Executor {
	return u.client
}

func (u UserRepository) getByNameOrEmail(ctx context.Context, f fn, query string, args ...any) (generated.IEntity, error) {
	row, err := f(ctx, u.client, query, args...)
	if err != nil {
		return nil, err
	}
	return marshalRow(row)
}

func marshalRow(row *sql.Row) (generated.IEntity, error) {
	var jsonData []byte
	err := row.Scan(&jsonData)
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
