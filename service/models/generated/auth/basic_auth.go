package auth

type BasicAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
