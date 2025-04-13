package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/TeddyCr/priceitt/service/errors"
	"github.com/TeddyCr/priceitt/service/infrastructure/jwt_secret"
	"github.com/TeddyCr/priceitt/service/models/types"
	repository "github.com/TeddyCr/priceitt/service/repository/database"
	"github.com/TeddyCr/priceitt/service/utils/jwt"
	"github.com/go-chi/render"

	goErrors "github.com/pkg/errors"
)

// JWTCtx is a middleware that validates the JWT token and adds the claims to the context
func JWTCtx(tokenService jwt_secret.ITokenService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if token == "" {
				err := render.Render(w, r, errors.ErrUnauthorized(goErrors.New("no token provided")))
				if err != nil {
					panic(err)
				}
				return
			}
			splitToken := strings.Split(token, "Bearer ")
			if len(splitToken) != 2 {
				err := render.Render(w, r, errors.ErrUnauthorized(goErrors.New("invalid token")))
				if err != nil {
					panic(err)
				}
				return
			}
			token = splitToken[1]
			claims, err := jwt.DecodeAndValidateJWT(token)
			if err != nil {
				err := render.Render(w, r, errors.ErrUnauthorized(err))
				if err != nil {
					panic(err)
				}
				return
			}
			userId, err := claims.GetSubject()
			if err != nil {
				err := render.Render(w, r, errors.ErrUnauthorized(err))
				if err != nil {
					panic(err)
				}
				return
			}
			decodedToken, err := jwt.DecodeJWT(token)
			if err != nil {
				err := render.Render(w, r, errors.ErrUnauthorized(err))
				if err != nil {
					panic(err)
				}
				return
			}
			isBlacklisted, err := tokenService.IsTokenBlacklisted(r.Context(), decodedToken, repository.QueryFilter{})
			if err != nil {
				err := render.Render(w, r, errors.ErrUnauthorized(err))
				if err != nil {
					panic(err)
				}
				return
			}
			if isBlacklisted {
				err := render.Render(w, r, errors.ErrUnauthorized(goErrors.New("Invalid token, token is blacklisted")))
				if err != nil {
					panic(err)
				}
				return
			}
			jwtContextValues := types.JWTContextValues{
				M: map[string]any{
					"claims": claims,
					"token":  string(decodedToken),
					"userId": userId,
				},
			}
			// TODO implement https://stackoverflow.com/questions/40891345/fix-should-not-use-basic-type-string-as-key-in-context-withvalue-golint
			ctx := context.WithValue(r.Context(), "jwtContextValues", jwtContextValues) //nolint:staticcheck
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
