package handlers

import (
	"log"
	"net/http"
	"os"
	"time"

	"go_mangasnake_api/api/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var jwtSecret = []byte(os.Getenv("SECRET_TOKEN"))

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to connect to Database:", err)
	}

	dsn := os.Getenv("DATABASE_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to Database:", err) // duplicaod a cima, mudar dps pra um arquivo de handlers/errors.go pra lidar com msgs
	}
	if err := DB.AutoMigrate(&model.Manga{}); err != nil {

		log.Fatal("failed to migrate schema:", err)
	}
}

func CreateManga(c *gin.Context) {
	var manga model.Manga

	if err := c.ShouldBindJSON(&manga); err != nil {
		model.ResponseJSON(c, http.StatusBadRequest, "invalid input", nil)
		return
	}
	DB.Create(&manga)
	model.ResponseJSON(c, http.StatusCreated, "Manga created successfully!!!", manga)
}

func GetManga(c *gin.Context) {
	var manga model.Manga
	if err := DB.First(&manga, c.Param("id")).Error; err != nil {
		model.ResponseJSON(c, http.StatusNotFound, "Manga not found :(", nil)
		return
	}
	model.ResponseJSON(c, http.StatusOK, "Manga retrieved successfully", manga)
}

func GetMangas(c *gin.Context) {
	var mangas []model.Manga
	DB.Find(&mangas)
	model.ResponseJSON(c, http.StatusOK, "Mangas retrieved successfully", mangas)
}

func UpdateManga(c *gin.Context) {
	var manga model.Manga
	if err := DB.First(&manga, c.Param("id")).Error; err != nil {
		model.ResponseJSON(c, http.StatusNotFound, "Manga not found :(", nil)
		return
	}

	if err := c.ShouldBindJSON(&manga); err != nil {
		model.ResponseJSON(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}
	DB.Save(&manga)
	model.ResponseJSON(c, http.StatusOK, "Mangas updated successfully", manga)
}

func DeleteManga(c *gin.Context) {
	var manga model.Manga
	if err := DB.Delete(&manga, c.Param("id")).Error; err != nil {
		model.ResponseJSON(c, http.StatusNotFound, "Manga not found :(", nil)
		return
	}
	model.ResponseJSON(c, http.StatusOK, "Manga successfully deleted", nil)
}

func GenerateJWT(c *gin.Context) {
	var loginRequest model.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		model.ResponseJSON(c, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}
	if loginRequest.Username != "admin" || loginRequest.Password != "password" {
		model.ResponseJSON(c, http.StatusUnauthorized, "Invalid credentials", nil)
		return
	}
	expirationTime := time.Now().Add(15 * time.Minute)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": expirationTime.Unix(),
	})
	// Sign the token
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		model.ResponseJSON(c, http.StatusInternalServerError, "Could not generate token", nil)
		return
	}
	model.ResponseJSON(c, http.StatusOK, "Token generated successfully", gin.H{"token": tokenString})
}
