package repository

const (
	InsertQuery = `INSERT INTO %s (json) VALUES ($1::jsonb)`
)
