package user_repository

const (
	GetByEmail = `SELECT json FROM %s WHERE %s AND email = $1`
)
