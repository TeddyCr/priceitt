package database

const (
	InsertQuery = `INSERT INTO %s (json) VALUES ($1)`
)