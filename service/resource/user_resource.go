package resource

import (
	"encoding/json"
	"net/http"

	"github.com/TeddyCr/priceitt/service/errors"
	"github.com/TeddyCr/priceitt/service/handler"
	"github.com/TeddyCr/priceitt/service/handler/user"
	"github.com/TeddyCr/priceitt/service/infrastructure/database"
	"github.com/TeddyCr/priceitt/service/infrastructure/jwt_secret"
	"github.com/TeddyCr/priceitt/service/middleware"
	authModels "github.com/TeddyCr/priceitt/service/models/generated/auth"
	"github.com/TeddyCr/priceitt/service/models/generated/createEntities"
	repository "github.com/TeddyCr/priceitt/service/repository/database"
	auth "github.com/TeddyCr/priceitt/service/repository/database/auth"
	usr "github.com/TeddyCr/priceitt/service/repository/database/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func NewUserResource(databasePersitence database.IPersistenceDatabase) IResource {
	databaseRepository := usr.NewUserRepository(databasePersitence)
	authRepository := auth.NewAuthRepository(databasePersitence)
	handler := user.NewUserHandler(databaseRepository, authRepository)
	return userResource{
		_user_handler: handler,
	}
}

type userResource struct {
	_user_handler handler.IHandler
}

func (ur userResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", ur.CreateUser)
	r.Post("/login", ur.Login)

	r.Group(func(r chi.Router) {
		r.Use(middleware.JWTCtx(jwt_secret.GetTokenServiceInstance()))
		r.Post("/logout", ur.Logout)
	})

	return r
}

func (ur userResource) CreateUser(w http.ResponseWriter, r *http.Request) {
	createUser := &createEntities.CreateUser{}
	if err := render.Bind(r, createUser); err != nil {
		err := render.Render(w, r, errors.ErrInvalidRequest(err))
		if err != nil {
			panic(err)
		}
		return
	}

	user, err := ur._user_handler.Create(r.Context(), createUser, *repository.NewQueryFilter(nil))
	if err != nil {
		err := render.Render(w, r, errors.ErrInternalServer(err))
		if err != nil {
			panic(err)
		}
		return
	}
	render.Status(r, http.StatusCreated)
	err = render.Render(w, r, user)
	if err != nil {
		panic(err)
	}
}

func (ur userResource) Login(w http.ResponseWriter, r *http.Request) {
	var authEncapsuler authModels.AuthEncapsulation
	err := json.NewDecoder(r.Body).Decode(&authEncapsuler)
	if err != nil {
		err := render.Render(w, r, errors.ErrInvalidRequest(err))
		if err != nil {
			panic(err)
		}
		return
	}
	tokens, err := ur._user_handler.(user.UserHandler).Login(r.Context(), authEncapsuler)
	if err != nil {
		err := render.Render(w, r, errors.ErrInternalServer(err))
		if err != nil {
			panic(err)
		}
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, tokens)
}

func (ur userResource) Logout(w http.ResponseWriter, r *http.Request) {
	_, err := ur._user_handler.(user.UserHandler).Logout(r.Context())
	if err != nil {
		err := render.Render(w, r, errors.ErrInternalServer(err))
		if err != nil {
			panic(err)
		}
		return
	}
	render.Status(r, http.StatusResetContent)
}
