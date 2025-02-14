package database

import (
	"go_mangasnake_api/api/model"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

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

	if err := DB.AutoMigrate(&model.User{}, &model.Manga{}, &model.List{}); err != nil {
		log.Fatal("Failed to migrate schema:", err)
	}

}
