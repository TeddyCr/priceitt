package database

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/TeddyCr/priceitt/service/models"
	"github.com/TeddyCr/priceitt/service/models/generated"
	"github.com/TeddyCr/priceitt/service/serializer"
	"github.com/jmoiron/sqlx"
	"go.uber.org/multierr"

	_ "github.com/lib/pq" // postgres driver
)

type Executor interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

func Connect(dbConfig models.DatabaseConfig) *sqlx.DB {
	db, err := sqlx.Connect(
		dbConfig.DriverClass,
		dbConfig.ConnectionString,
	)
	if err != nil {
		log.Fatalf("Error connecting to db: #%v", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db
}

func PerformEntityQuery(ctx context.Context, db Executor, query string, entity generated.IEntity) error {
	jsonString, err := serializer.JsonToString(entity)
	if err != nil {
		return err
	}
	_, err = db.ExecContext(ctx, query, jsonString)
	if err != nil {
		return err
	}
	return nil
}

func PerformEntityQueryTx(ctx context.Context, tx *sql.Tx, query string, entity generated.IEntity) error {
	jsonString, err := serializer.JsonToString(entity)
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, query, jsonString)
	if err != nil {
		return err
	}
	return nil
}

func PerformSelectScalarQueryTx(ctx context.Context, db *sql.Tx, query string, args ...any) (*sql.Row, error) {
	row := db.QueryRowContext(ctx, query, args...)
	if row.Err() != nil {
		return nil, row.Err()
	}
	return row, nil
}

func PerformSelectScalarQuery(ctx context.Context, db Executor, query string, args ...any) (*sql.Row, error) {
	row := db.QueryRowContext(ctx, query, args...)
	if row.Err() != nil {
		return nil, row.Err()
	}
	return row, nil
}

func FinishTx(err error, tx *sql.Tx) error {
	if err != nil {
		if rollBackErr := tx.Rollback(); rollBackErr != nil {
			return multierr.Combine(rollBackErr, err)
		}
	} else {
		if commitErr := tx.Commit(); commitErr != nil {
			return multierr.Combine(commitErr, err)
		}
	}
	return err
}

func RunInTx(ctx context.Context, db *sqlx.DB, fn func(tx *sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err := FinishTx(err, tx); err != nil {
			log.Fatalf("Error finishing tx: #%v", err)
		}
	}()

	err = fn(tx)
	if err != nil {
		return err
	}
	return nil
}
