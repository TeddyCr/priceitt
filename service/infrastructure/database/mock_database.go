package database

import (
	"context"
	"database/sql"

	"github.com/TeddyCr/priceitt/service/models"
	"github.com/TeddyCr/priceitt/service/utils/database"
)

type MockExecutor struct {
}

func (m MockExecutor) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return nil, nil
}

func (m MockExecutor) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return nil, nil
}

func (m MockExecutor) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return nil
}

type MockDatabase struct {
	database.Executor
}

func (m MockDatabase) Initialize(config models.DatabaseConfig) (IPersistenceDatabase, error) {
	return m, nil
}

func (m MockDatabase) GetClient() database.Executor {
	return m.Executor
}
