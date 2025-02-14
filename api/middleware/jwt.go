package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTU(userID uint) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := jwt.MapClaims{
		"exp": expirationTime.Unix(),
	}

	if userID == 0 {
		// Se o ID for 0, consideramos como admin
		claims["admin"] = true
	} else {
		claims["user_id"] = userID
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
