package auth_repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/TeddyCr/priceitt/service/infrastructure/database"
	"github.com/TeddyCr/priceitt/service/models/generated"
	"github.com/TeddyCr/priceitt/service/models/generated/entities"
	"github.com/TeddyCr/priceitt/service/models/types"
	repository "github.com/TeddyCr/priceitt/service/repository/database"
	utilDB "github.com/TeddyCr/priceitt/service/utils/database"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func NewAuthRepository(dbContext database.IPersistenceDatabase) *AuthRepository {
	return &AuthRepository{dbContext: dbContext, client: dbContext.GetClient()}
}

// AuthRepository is a struct that defines the methods that an auth repository should implement
type AuthRepository struct {
	dbContext database.IPersistenceDatabase
	client    utilDB.Executor
}

func (a *AuthRepository) Create(ctx context.Context, token generated.IEntity) error {
	query := fmt.Sprintf(repository.InsertQuery, "tokens")
	err := utilDB.PerformEntityQuery(ctx, a.client, query, token)
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthRepository) CreateBlacklistToken(ctx context.Context, token generated.IEntity) error {
	query := fmt.Sprintf(repository.InsertQuery, "token_blacklist")
	err := utilDB.PerformEntityQuery(ctx, a.client, query, token)
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthRepository) GetById(ctx context.Context, id string, filter repository.QueryFilter) (generated.IEntity, error) { // (entities.User, error) {
	return nil, nil
}

func (a *AuthRepository) GetByName(ctx context.Context, name string, filter repository.QueryFilter) (generated.IEntity, error) { // (entities.User, error) {
	return nil, nil
}

func (a *AuthRepository) UpdateById(ctx context.Context, id string, entity generated.IEntity, filter repository.QueryFilter) error {
	return nil
}

func (a *AuthRepository) UpdateByName(ctx context.Context, name string, entity generated.IEntity, filter repository.QueryFilter) error {
	return nil
}

func (a *AuthRepository) DeleteById(ctx context.Context, id string, filter repository.QueryFilter) error {
	query := fmt.Sprintf(repository.DeleteById, "tokens", filter.String())
	filter.Filter["id"] = id
	_, err := a.client.ExecContext(ctx, query, filter.Args()...)
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthRepository) DeleteByName(ctx context.Context, name string, filter repository.QueryFilter) error {
	query := fmt.Sprintf(repository.DeleteByName, "tokens", filter.String())
	filter.Filter["name"] = name
	_, err := a.client.ExecContext(ctx, query, filter.Args()...)
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthRepository) List(ctx context.Context, filter repository.QueryFilter) ([]generated.IEntity, error) { // ([]entities.User, error) {
	return nil, nil
}

func (a *AuthRepository) GetClient() utilDB.Executor {
	return a.client
}

func (a *AuthRepository) Logout(ctx context.Context, token string, user entities.User) error {
	db := a.client.(*sqlx.DB)
	err := utilDB.RunInTx(ctx, db, func(tx *sql.Tx) error {
		newFilter := repository.NewQueryFilter(map[string]string{
			"userId": user.ID.String(),
		})
		query := fmt.Sprintf(repository.GetByName, "tokens", newFilter.String())
		// TODO: handle this better. We need to recreate the filter because
		// name needs to be the first parameter
		newFilter = repository.NewQueryFilter(map[string]string{
			"name": "refresh",
			"userId": user.ID.String(),
		})
		row, err := utilDB.PerformSelectScalarQueryTx(ctx, tx, query, newFilter.Args()...)
		if err != nil {
			return err
		}
		marshaledRow, err := marshalRow(row)
		if err != nil {
			return err
		}
		authRepository := AuthRepository{a.dbContext, tx}
		authRepository.DeleteById(ctx, marshaledRow.GetID().String(), *repository.NewQueryFilter(nil))
		authRepository.CreateBlacklistToken(ctx, marshaledRow)
		authRepository.CreateBlacklistToken(ctx, &entities.JWToken{
			ID: uuid.New(),
			IP: "",
			Token: token,
			TokenType: types.TokenType(types.AccessToken).String(),
			UserID: user.ID,
		})

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func marshalRow(row *sql.Row) (generated.IEntity, error) {
	var jsonData []byte
	err := row.Scan(&jsonData)
	if err != nil {
		return nil, err
	}
	var token entities.JWToken
	err = json.Unmarshal(jsonData, &token)
	if err != nil {
		return nil, err
	}
	return &token, nil
}
