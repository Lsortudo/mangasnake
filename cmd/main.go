package main

import (
	"go_mangasnake_api/api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	handlers.InitDB()
	r := gin.Default()

	r.POST("/manga", handlers.CreateManga)
	r.GET("/mangas", handlers.GetMangas)
	r.GET("/manga/:id", handlers.GetManga)
	r.PUT("/manga/:id", handlers.UpdateManga)
	r.DELETE("/manga/:id", handlers.DeleteManga)

	r.Run(":8080")

}
