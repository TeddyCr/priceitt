package auth

type Basic struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	Type            string `json:"type"`
}

func (b Basic) GetAuthType() string {
	return b.Type
}

func (b Basic) GetPassword() string {
	return b.Password
}
