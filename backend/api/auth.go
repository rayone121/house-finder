package api

import (
	"fmt"
	"net/http"
	"os"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/rayone121/reBank/backend/storage"
	"github.com/rayone121/reBank/backend/types"
)

func permissionDenied(w http.ResponseWriter) {
	writeJSON(w, http.StatusForbidden, apiError{Error: "permission denied"})
}

func withJWTAuth(next http.HandlerFunc, s storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("x-auth-token")
		token, err := validateJWT(tokenString)
		if err != nil {
			permissionDenied(w)
			return
		}

		if !token.Valid {
			permissionDenied(w)
			return
		}
		userID, err := getID(r)
		if err != nil {
			permissionDenied(w)
			return
		}

		account, err := s.GetAccountByID(userID)
		if err != nil {
			permissionDenied(w)
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		if account.Number != int64(claims["accountNumber"].(float64)) {
			permissionDenied(w)
			return
		}

		next(w, r)
	}
}

// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50SUQiOjAsImV4cGlyZUF0IjoxNTAwMH0.ycyUe2i973kp1c99q2WHfR6kp42_lBtCz1v1V27H4fQ

func createJWT(account *types.Account) (string, error) {
	claims := &jwt.MapClaims{
		"expireAt":      1500,
		"accountNumber": account.Number,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("SECRET_KEY")
	return token.SignedString([]byte(secret))
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
