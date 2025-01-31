package main

import (
	"go_mangasnake_api/api/handlers"
	"go_mangasnake_api/api/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	handlers.InitDB()
	r := gin.Default()

	// rotas publicas
	r.POST("/token", handlers.GenerateJWT)

	// rotas protegidas
	protected := r.Group("/", middleware.JWTAuthMiddleware())
	{
		protected.POST("/manga", handlers.CreateManga)
		protected.GET("/mangas", handlers.GetMangas)
		protected.GET("/manga/:id", handlers.GetManga)
		protected.PUT("/manga/:id", handlers.UpdateManga)
		protected.DELETE("/manga/:id", handlers.DeleteManga)

	}

	r.Run(":8080")

}
