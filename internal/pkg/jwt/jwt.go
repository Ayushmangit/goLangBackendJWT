package jwt

import (
	"fmt"
	"time"

	"github.com/Ayushmangit/goLangBackendJWT/internal/env"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(env.GetString("SECRET", "asdagadhgfsadhfadf312"))

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(userId, email string) (string, error) {
	claims := Claims{
		UserID: userId,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	fmt.Println("GENERATING TOKEN USER ID:", claims.UserID)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
