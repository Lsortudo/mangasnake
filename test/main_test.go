package test

import (
	"bytes"
	"encoding/json"
	"go_mangasnake_api/api/handlers"
	"go_mangasnake_api/api/model"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var jwtSecret = []byte(os.Getenv("SECRET_TOKEN"))

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

func TestGetManga(t *testing.T) {
	setupTestDB()
	manga := addManga()
	router := gin.Default()
	router.GET("/manga/:id", handlers.GetManga)

	req, _ := http.NewRequest("GET", "/manga/"+strconv.Itoa(int(manga.ID)), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}

	var response model.JsonResponse
	json.NewDecoder(w.Body).Decode(&response)

	if response.Data == nil || response.Data.(map[string]interface{})["id"] != float64(manga.ID) {
		t.Errorf("Expected manga ID %d, got nil or wrong ID", manga.ID)
	}
}

func TestUpdateManga(t *testing.T) {
	setupTestDB()
	manga := addManga()
	router := gin.Default()
	router.PUT("/manga/:id", handlers.UpdateManga)

	updateManga := model.Manga{
		Title: "Teste manga update", Author: "LeoSan", Year: 2026,
	}
	jsonValue, _ := json.Marshal(updateManga)

	req, _ := http.NewRequest("PUT", "/manga/"+strconv.Itoa(int(manga.ID)), bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}

	var response model.JsonResponse
	json.NewDecoder(w.Body).Decode(&response)

	if response.Data == nil || response.Data.(map[string]interface{})["title"] != "Teste manga update" {
		t.Errorf("Expected updated manga title 'Teste manga update', got %v", response.Data)

	}
}

func TestDeleteManga(t *testing.T) {
	setupTestDB()
	manga := addManga()
	router := gin.Default()
	router.DELETE("/manga/:id", handlers.DeleteManga)

	req, _ := http.NewRequest("DELETE", "/manga/"+strconv.Itoa(int(manga.ID)), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}

	var response model.JsonResponse
	json.NewDecoder(w.Body).Decode(&response)

	if response.Message != "Manga successfully deleted" {
		t.Errorf("Expected delete message 'Manga successfully deleted', got %v", response.Message)
	}

	var deletedManga model.Manga
	result := handlers.DB.First(&deletedManga, manga.ID)
	if result.Error == nil {
		t.Errorf("Expected manga to be deleted, but it still exists")

	}
}

func generateValidToken() string {
	expirationTime := time.Now().Add(15 * time.Minute)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": expirationTime.Unix(),
	})
	tokenString, _ := token.SignedString(jwtSecret)
	return tokenString
}

func TestGenerateJWT(t *testing.T) {
	router := gin.Default()
	router.POST("/token", handlers.GenerateJWT)

	loginRequest := map[string]string{
		"username": "admin",
		"password": "password",
	}

	jsonValue, _ := json.Marshal(loginRequest)
	req, _ := http.NewRequest("POST", "/token", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}

	var response model.JsonResponse
	json.NewDecoder(w.Body).Decode(&response)

	if response.Data == nil || response.Data.(map[string]interface{})["token"] == "" {
		t.Errorf("Expected token in response, got nil or empty")
	}
}
