package auth

type BaseAuthMechanism interface {
	GetAuthType() string
}

type Basic struct {
	Password string `json:"password"`
	Type string `json:"type"`
}

func (b Basic) GetAuthType() string {
	return b.Type
}
