package middleware

import (
	"context"
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/TeddyCr/priceitt/service/application/user"
	"github.com/TeddyCr/priceitt/service/models/types"
)

// JWTCtx is a middleware that validates the JWT token and adds the claims to the context
func JWTCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unauthorized: No token provided", http.StatusUnauthorized)
			return
		}
		splitToken := strings.Split(token, "Bearer ")
		if len(splitToken) != 2 {
			http.Error(w, "Unauthorized: Invalid token format", http.StatusUnauthorized)
			return
		}
		token = splitToken[1]
		decodedToken, err := base64.StdEncoding.DecodeString(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		claims, err := user.ValidateJWT(string(decodedToken))
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		userId, err := claims.GetSubject()
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		jwtContextValues := types.JWTContextValues{
			M: map[string]any{
				"claims": claims,
				"token":  string(decodedToken),
				"userId": userId,
			},
		}
		ctx := context.WithValue(r.Context(), "jwtContextValues", jwtContextValues)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
	