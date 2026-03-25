package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Ayushmangit/goLangBackendJWT/internal/env"
	jwtpkg "github.com/Ayushmangit/goLangBackendJWT/internal/pkg/jwt"
	"github.com/Ayushmangit/goLangBackendJWT/internal/types"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(env.GetString("SECRET", "asdagadhgfsadhfadf312"))

type contextKey string

const UserKey contextKey = "user"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		claims := &jwtpkg.Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (any, error) {
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		authUser := &types.AuthUser{
			ID:    claims.UserID,
			Email: claims.Email,
		}

		ctx := context.WithValue(r.Context(), UserKey, authUser)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
