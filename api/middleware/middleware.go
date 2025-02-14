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

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			model.ResponseJSON(ctx, http.StatusUnauthorized, "Authorization token required", nil)
			ctx.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})

		if err != nil || token == nil {
			model.ResponseJSON(ctx, http.StatusUnauthorized, "Invalid token", nil)
			ctx.Abort()
			return
		}

		// Verifica se o token contém a chave "admin"
		claims, ok := token.Claims.(jwt.MapClaims)
		isAdmin, hasAdminKey := claims["admin"].(bool)

		if !ok || !hasAdminKey || !isAdmin {
			model.ResponseJSON(ctx, http.StatusForbidden, "Access denied", nil)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

// func GenerateJWT(c *gin.Context) {
// 	var loginRequest model.LoginRequest
// 	if err := c.ShouldBindJSON(&loginRequest); err != nil {
// 		model.ResponseJSON(c, http.StatusBadRequest, "Invalid request payload", nil)
// 		return
// 	}
// 	if loginRequest.Email != "admin@teste.com" || loginRequest.Password != "adminop1" {
// 		model.ResponseJSON(c, http.StatusUnauthorized, "Invalid credentials", nil)
// 		return
// 	}
// 	expirationTime := time.Now().Add(15 * time.Minute)
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"exp": expirationTime.Unix(),
// 	})
// 	// Sign the token
// 	tokenString, err := token.SignedString(jwtSecret)
// 	if err != nil {
// 		model.ResponseJSON(c, http.StatusInternalServerError, "Could not generate token", nil)
// 		return
// 	}
// 	model.ResponseJSON(c, http.StatusOK, "Token generated successfully", gin.H{"token": tokenString})
// }

// colocar depois apenas uma funcao GenerateJWT, e pra generateJWT padrao que lista mangas fazer de algum jeito que só sirva com o login admin admin
