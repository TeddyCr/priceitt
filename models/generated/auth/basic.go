package auth

type Basic struct {
	Password string `json:"password"`
	Type string `json:"type"`
}
