package resource

import (
	"net/http"

	errorsService "github.com/TeddyCr/priceitt/service/errors"
	"github.com/TeddyCr/priceitt/service/handler"
	"github.com/TeddyCr/priceitt/service/handler/auth"
	"github.com/TeddyCr/priceitt/service/infrastructure/database"
	"github.com/TeddyCr/priceitt/service/infrastructure/jwt_secret"
	"github.com/TeddyCr/priceitt/service/middleware"
	"github.com/TeddyCr/priceitt/service/models/types"
	authRepository "github.com/TeddyCr/priceitt/service/repository/database/auth"
	"github.com/TeddyCr/priceitt/service/utils/jwt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func NewAuthResource(databasePersitence database.IPersistenceDatabase) IResource {
	databaseRepository := authRepository.NewAuthRepository(databasePersitence)
	jwt_secret.InitializeTokenService(*databaseRepository)
	handler := auth.NewAuthHandler(databaseRepository)
	return authResource{
		_auth_handler: handler,
	}
}

type authResource struct {
	_auth_handler handler.IHandler
}

func (ar authResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.JWTCtx(jwt_secret.GetTokenServiceInstance()))
	r.Post("/refresh", ar.RefreshAccessToken)

	return r
}

func (ar authResource) RefreshAccessToken(w http.ResponseWriter, r *http.Request) {
	jwtContextValues := r.Context().Value("jwtContextValues").(types.JWTContextValues)
	xRefreshToken, err := r.Cookie("X-Refresh-Token")
	if err != nil {
		err := render.Render(w, r, errorsService.ErrInternalServer(err))
		if err != nil {
			panic(err)
		}
	}
	refreshToken := xRefreshToken.Value
	userId := jwtContextValues.Get("userId").(string)

	_, err = jwt.DecodeAndValidateJWT(refreshToken)
	if err != nil {
		err := render.Render(w, r, errorsService.ErrInternalServer(err))
		if err != nil {
			panic(err)
		}
	}
	accessToken, err := ar._auth_handler.(auth.AuthHandler).CreateAccessToken(r.Context(), userId)
	if err != nil {
		err := render.Render(w, r, errorsService.ErrInternalServer(err))
		if err != nil {
			panic(err)
		}
		return
	}
	err = render.Render(w, r, accessToken)
	if err != nil {
		panic(err)
	}
}
