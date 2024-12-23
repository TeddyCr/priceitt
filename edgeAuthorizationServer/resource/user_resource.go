package resource

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/TeddyCr/priceitt/edgeAuthorizationServer/application"
	"github.com/TeddyCr/priceitt/edgeAuthorizationServer/application/user"
	"github.com/TeddyCr/priceitt/edgeAuthorizationServer/infrastructure/database"
	usr "github.com/TeddyCr/priceitt/edgeAuthorizationServer/repository/database/user"
)

func NewUserResource(databasePersitence database.IPersistenceDatabase) IUserResource {
	databaseRepository := usr.NewUserRepository(databasePersitence)
	handler := user.NewUserHandler(databaseRepository)
	return userResource{
		_user_handler: handler,
	}
}

type userResource struct {
	_user_handler application.IHandler
}

func (ur userResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", ur.CreateUser)

	return r
}

func (ur userResource) CreateUser(w http.ResponseWriter, r *http.Request) {
	ur._user_handler.Create(r.Context(), nil)
}
