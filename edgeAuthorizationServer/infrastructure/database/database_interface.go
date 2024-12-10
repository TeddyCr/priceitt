package database

import (
	"github.com/TeddyCr/priceitt/models"
	"github.com/jmoiron/sqlx"
)

type IPersistenceDatabase interface {
	Initialize(config models.DatabaseConfig) (IPersistenceDatabase, error)
	GetClient() *sqlx.DB
}