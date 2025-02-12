package middleware

import (
	"fmt"
	"go_mangasnake_api/api/model"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("SECRET_TOKEN"))

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			model.ResponseJSON(ctx, http.StatusUnauthorized, "Authorization token required", nil)
			ctx.Abort()
			return
		}

		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})

		if err != nil {
			model.ResponseJSON(ctx, http.StatusUnauthorized, "invalid token", nil)
			ctx.Abort()
			return
		}
		// Token valido, proximo handler
		ctx.Next()
	}
}

func GenerateJWTU(userID uint) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     expirationTime.Unix(),
	})
	return token.SignedString(jwtSecret)
}
