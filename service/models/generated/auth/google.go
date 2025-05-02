package auth

type Google struct {
	Type     string `json:"type"`
	IdToken  string `json:"idToken"`
	Audience string `json:"audience"`
}

func (g Google) GetAuthType() string {
	return g.Type
}

func (g Google) GetIdToken() string {
	return g.IdToken
}

func (g Google) GetAudience() string {
	return g.Audience
}
