package middleware

import (
	"context"
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/TeddyCr/priceitt/service/errors"
	"github.com/TeddyCr/priceitt/service/models/types"
	"github.com/TeddyCr/priceitt/service/utils/jwt"
	"github.com/go-chi/render"
)

// JWTCtx is a middleware that validates the JWT token and adds the claims to the context
func JWTCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			err := render.Render(w, r, errors.ErrUnauthorized())
			if err != nil {
				panic(err)
			}
			return
		}
		splitToken := strings.Split(token, "Bearer ")
		if len(splitToken) != 2 {
			err := render.Render(w, r, errors.ErrUnauthorized())
			if err != nil {
				panic(err)
			}
			return
		}
		token = splitToken[1]
		decodedToken, err := base64.StdEncoding.DecodeString(token)
		if err != nil {
			err := render.Render(w, r, errors.ErrUnauthorized())
			if err != nil {
				panic(err)
			}
			return
		}
		claims, err := jwt.ValidateJWT(string(decodedToken))
		if err != nil {
			err := render.Render(w, r, errors.ErrUnauthorized())
			if err != nil {
				panic(err)
			}
			return
		}
		userId, err := claims.GetSubject()
		if err != nil {
			err := render.Render(w, r, errors.ErrUnauthorized())
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
