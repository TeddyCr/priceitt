package database

import (
	"github.com/TeddyCr/priceitt/service/models"
	"github.com/TeddyCr/priceitt/service/utils/database"
)

type IPersistenceDatabase interface {
	Initialize(config models.DatabaseConfig) (IPersistenceDatabase, error)
	GetClient() database.Executor
}
