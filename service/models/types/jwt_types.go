package types

type TokenType int8

const (
	RefreshToken TokenType = iota
	AccessToken
)

func (t TokenType) String() string {
	switch t {
	case RefreshToken:
		return "refresh"
	case AccessToken:
		return "access"
	default:
		return "unknown"
	}
}