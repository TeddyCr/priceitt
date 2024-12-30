package generated

import "net/http"

type ICreateEntity interface {
	GetName() string
	Bind(r *http.Request) error
	Render(w http.ResponseWriter, r *http.Request) error
}
