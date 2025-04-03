package repository

const (
	InsertQuery = `INSERT INTO %s (json) VALUES ($1::jsonb)`
)

const (
	GetByName = `SELECT json FROM %s WHERE name = $1`
)
