package resource

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"priceitt.xyz/edgeAuthorizationServer/application"
	"priceitt.xyz/edgeAuthorizationServer/application/user"
	"priceitt.xyz/edgeAuthorizationServer/infrastructure/database"
	usr "priceitt.xyz/edgeAuthorizationServer/repository/database/user"
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
