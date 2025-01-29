package test

import (
	"bytes"
	"encoding/json"
	"go_mangasnake_api/api/handlers"
	"go_mangasnake_api/api/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() {
	var err error

	handlers.DB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect test databasea")
	}
	handlers.DB.AutoMigrate(&model.Manga{})
}

func addManga() model.Manga {
	manga := model.Manga{Title: "Tower of Gado", Author: "SIIIIIU", Year: 2010}
	handlers.DB.Create(&manga)
	return manga
}

func TestCreateManga(t *testing.T) {
	setupTestDB()
	router := gin.Default()

	router.POST("/manga", handlers.CreateManga)

	manga := model.Manga{
		Title: "Manga Func Create", Author: "LeoS", Year: 2025,
	}

	jsonValue, _ := json.Marshal(manga)
	req, _ := http.NewRequest("POST", "/manga", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, status)
	}
	var response model.JsonResponse
	json.NewDecoder(w.Body).Decode(&response)

	if response.Data == nil {
		t.Errorf("Expected manga data, got nil")
	}
}

func TestGetMangas(t *testing.T) {
	setupTestDB()
	addManga()
	router := gin.Default()

	router.GET("/mangas", handlers.GetMangas)
	req, _ := http.NewRequest("GET", "/mangas", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}

	var response model.JsonResponse
	json.NewDecoder(w.Body).Decode(&response)

	if response.Data == nil {
		t.Errorf("Expected non-empty mangas list")
	}
}
