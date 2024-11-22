package middleware

import (
	"context"
	"net/http"
	"orders/internal/configs"
	"orders/pkg/jwt"
	"strings"
)

func writeUnauthed(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
}

func IsAuthed(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authedHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authedHeader, "Bearer ") {
			writeUnauthed(w)
			return
		}
		token := strings.TrimPrefix(authedHeader, "Bearer ")
		isValid, data := jwt.NewJWT(config.JWTSecret).Parse(token)
		if !isValid {
			writeUnauthed(w)
			return
		}
		ctx := context.WithValue(r.Context(), "id", data.Id)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
