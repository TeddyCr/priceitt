package database

import (
	"github.com/TeddyCr/priceitt/service/models"
	"github.com/jmoiron/sqlx"
)

type IPersistenceDatabase interface {
	Initialize(config models.DatabaseConfig) (IPersistenceDatabase, error)
	GetClient() *sqlx.DB
}
