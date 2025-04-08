package postgres

import (
	"github.com/TeddyCr/priceitt/service/infrastructure/database"
	"github.com/TeddyCr/priceitt/service/models"
	dbUtil "github.com/TeddyCr/priceitt/service/utils/database"
	_ "github.com/lib/pq"
)

type PersistencePostgres struct {
	client dbUtil.Executor
}

func (PersistencePostgres) Initialize(config models.DatabaseConfig) (database.IPersistenceDatabase, error) {
	client := dbUtil.Connect(config)
	return PersistencePostgres{client: client}, nil
}

func (p PersistencePostgres) GetClient() dbUtil.Executor {
	return p.client
}
