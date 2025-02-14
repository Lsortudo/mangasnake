package handlers

import (
	"net/http"

	"go_mangasnake_api/api/database"
	"go_mangasnake_api/api/middleware"
	"go_mangasnake_api/api/model"

	"github.com/gin-gonic/gin"
)

func CreateManga(c *gin.Context) {
	var manga model.Manga

	if err := c.ShouldBindJSON(&manga); err != nil {
		model.ResponseJSON(c, http.StatusBadRequest, "invalid input", nil)
		return
	}
	database.DB.Create(&manga)
	model.ResponseJSON(c, http.StatusCreated, "Manga created successfully!!!", manga)
}

func GetManga(c *gin.Context) {
	var manga model.Manga
	if err := database.DB.First(&manga, c.Param("id")).Error; err != nil {
		model.ResponseJSON(c, http.StatusNotFound, "Manga not found :(", nil)
		return
	}
	model.ResponseJSON(c, http.StatusOK, "Manga retrieved successfully", manga)
}

func GetMangas(c *gin.Context) {
	var mangas []model.Manga
	database.DB.Find(&mangas)
	model.ResponseJSON(c, http.StatusOK, "Mangas retrieved successfully", mangas)
}

func UpdateManga(c *gin.Context) {
	var manga model.Manga
	if err := database.DB.First(&manga, c.Param("id")).Error; err != nil {
		model.ResponseJSON(c, http.StatusNotFound, "Manga not found :(", nil)
		return
	}

	if err := c.ShouldBindJSON(&manga); err != nil {
		model.ResponseJSON(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}
	database.DB.Save(&manga)
	model.ResponseJSON(c, http.StatusOK, "Mangas updated successfully", manga)
}

func DeleteManga(c *gin.Context) {
	var manga model.Manga
	if err := database.DB.Delete(&manga, c.Param("id")).Error; err != nil {
		model.ResponseJSON(c, http.StatusNotFound, "Manga not found :(", nil)
		return
	}
	model.ResponseJSON(c, http.StatusOK, "Manga successfully deleted", nil)
}

func GenerateAdminToken(c *gin.Context) {
	var loginRequest model.LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		model.ResponseJSON(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}

	// Verifica se é o admin
	if loginRequest.Email != "admin@teste.com" || loginRequest.Password != "adminop1" {
		model.ResponseJSON(c, http.StatusUnauthorized, "Invalid credentials", nil)
		return
	}

	// Gera o token
	token, err := middleware.GenerateJWTU(0) // Pode usar um ID fictício para admin
	if err != nil {
		model.ResponseJSON(c, http.StatusInternalServerError, "Could not generate token", nil)
		return
	}

	model.ResponseJSON(c, http.StatusOK, "Token generated", gin.H{"token": token})
}
