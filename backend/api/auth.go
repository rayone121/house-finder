package api

import (
	"fmt"
	"net/http"
	"os"

	jwt "github.com/golang-jwt/jwt/v5"
)

func withJWTAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("x-auth-token")
		_, err := validateJWT(tokenString)
		if err != nil {
			writeJSON(w, http.StatusForbidden, apiError{Error: "invalid token"})
			return
		}

		next(w, r)
	}
}

func validateJWT(token string) (*jwt.Token, error) {
	secret := os.Getenv("SECRET_KEY")
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
}
