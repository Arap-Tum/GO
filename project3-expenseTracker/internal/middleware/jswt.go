package middleware

import (
	"context"
	"net/http"
	"strings"

	"expenseTracker/internal/utils"
)

type ContextKey string

const UserIDKey ContextKey = "user_id"

func JWTMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			auth := r.Header.Get("Authorization")

			if auth == "" {
				http.Error(
					w,
					"missing authorization",
					http.StatusUnauthorized,
				)
				return
			}

			parts := strings.Split(auth, " ")

			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(
					w,
					"invalid bearer token",
					http.StatusUnauthorized,
				)
				return
			}

			claims, err := utils.ValidateJWT(parts[1])

			if err != nil {
				http.Error(
					w,
					"unauthorized",
					http.StatusUnauthorized,
				)
				return
			}

			ctx := context.WithValue(
				r.Context(),
				UserIDKey,
				claims.UserID,
			)

			next.ServeHTTP(
				w,
				r.WithContext(ctx),
			)
		},
	)
}
