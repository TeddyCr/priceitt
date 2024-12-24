package resource

import (
	"net/http"

	"github.com/TeddyCr/priceitt/edgeAuthorizationServer/application"
	"github.com/TeddyCr/priceitt/edgeAuthorizationServer/errors"
	"github.com/TeddyCr/priceitt/edgeAuthorizationServer/application/user"
	"github.com/TeddyCr/priceitt/edgeAuthorizationServer/infrastructure/database"
	usr "github.com/TeddyCr/priceitt/edgeAuthorizationServer/repository/database/user"
	"github.com/TeddyCr/priceitt/models/generated/createEntities"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
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
	createUser := &createEntities.CreateUser{}
	if err := render.Bind(r, createUser); err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}

	user, err := ur._user_handler.Create(r.Context(), createUser)
	if err != nil {
		render.Render(w, r, errors.ErrInternalServer(err))
		return
	}
	render.Status(r, http.StatusCreated)
	render.Render(w, r, user)
	return
}
