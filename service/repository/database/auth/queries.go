package auth_repository

const (
	GetBlacklistToken = `SELECT json FROM token_blacklist WHERE %s AND token = $1`
)
