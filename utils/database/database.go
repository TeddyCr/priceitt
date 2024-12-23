package database

import (
	"context"
	"log"

	"github.com/TeddyCr/priceitt/models"
	"github.com/TeddyCr/priceitt/models/generated"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"           // postgres driver
)

func Connect(dbConfig models.DatabaseConfig) *sqlx.DB {
	db, err := sqlx.Connect(
		dbConfig.DriverClass,
		dbConfig.ConnectionString,
	)
	if err != nil {
		log.Fatalf("Error connecting to db: #%v", err)
	}
	return db
}

func PerformEntityQuery(ctx context.Context, db *sqlx.DB, query string, entity generated.IEntity) error {
	conn, err := db.Conn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()
	jsonString, err := entity.JsonToString()
	if err != nil {
		return err
	}
	_, err = conn.ExecContext(ctx, query, jsonString)
	if err != nil {
		return err
	}
	return nil
}
