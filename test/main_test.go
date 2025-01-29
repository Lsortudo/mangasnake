package test

import (
	"go_mangasnake_api/api/handlers"
	"go_mangasnake_api/api/model"

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
