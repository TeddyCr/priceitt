package user_repository

const (
	GetByEmail = `SELECT json FROM %s WHERE email = $1`
)
