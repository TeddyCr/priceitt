package postgres

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mitchellh/mapstructure"
	"github.com/TeddyCr/priceitt/models"
	"github.com/TeddyCr/priceitt/utils/database"
)

type PersistencePostgres struct {
	client *sqlx.DB
}

func (PersistencePostgres) Initialize(config map[string]interface{}) (*PersistencePostgres, error) {
	var dbConfig models.DatabaseConfig
	err := mapstructure.Decode(config, &dbConfig)
	if err != nil {
		log.Panic(err)
	}

	client := database.Connect(dbConfig)
	return &PersistencePostgres{client: client}, nil
}

func (p *PersistencePostgres) GetClient() *sqlx.DB {
	return p.client
}
