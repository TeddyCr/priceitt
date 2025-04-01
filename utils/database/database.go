package database

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/TeddyCr/priceitt/edgeAuthorizationServer/serializer"
	"github.com/TeddyCr/priceitt/models"
	"github.com/TeddyCr/priceitt/models/generated"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq" // postgres driver
)

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

func PerformEntityQuery(ctx context.Context, db *sqlx.DB, query string, entity generated.IEntity) error {
	conn, err := db.Conn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()
	jsonString, err := serializer.JsonToString(entity)
	if err != nil {
		return err
	}
	_, err = conn.ExecContext(ctx, query, jsonString)
	if err != nil {
		return err
	}
	return nil
}

func PerformSelectScalarQuery(ctx context.Context, db *sqlx.DB, query string, name string) (*sql.Row, error) {
	row := db.QueryRowContext(ctx, query, name)
	if row.Err() != nil {
		return nil, row.Err()
	}
	return row, nil
}
