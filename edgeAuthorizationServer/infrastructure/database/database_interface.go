package database

import (
	"github.com/jmoiron/sqlx"
)

type IPersistenceDatabase interface {
	Initialize(config map[string]interface{}) (*IPersistenceDatabase, error)
	GetClient() *sqlx.DB
}