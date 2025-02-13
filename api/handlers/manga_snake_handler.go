package handlers

import (
	"net/http"

	"go_mangasnake_api/api/model"

	"github.com/gin-gonic/gin"
)

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
