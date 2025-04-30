package auth

import "encoding/json"

type BaseAuthMechanism interface {
	GetAuthType() string
}

type AuthEncapsulation struct {
	Type string `json:"type"`
	Username string `json:"username"`
	Data json.RawMessage `json:"data"`
}

func (a AuthEncapsulation) GetAuthType() string {
	return a.Type
}

func (a AuthEncapsulation) GetData() json.RawMessage {
	return a.Data
}

func (a AuthEncapsulation) GetUsername() string {
	return a.Username
}
