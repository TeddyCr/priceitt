package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	"priceitt.xyz/models"
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