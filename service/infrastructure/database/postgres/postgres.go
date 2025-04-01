package postgres

import (
	"github.com/TeddyCr/priceitt/service/infrastructure/database"
	"github.com/TeddyCr/priceitt/service/models"
	dbUtil "github.com/TeddyCr/priceitt/service/utils/database"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PersistencePostgres struct {
	client *sqlx.DB
}

func (PersistencePostgres) Initialize(config models.DatabaseConfig) (database.IPersistenceDatabase, error) {
	client := dbUtil.Connect(config)
	return PersistencePostgres{client: client}, nil
}

func (p PersistencePostgres) GetClient() *sqlx.DB {
	return p.client
}
