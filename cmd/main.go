package main

import (
	"go_mangasnake_api/api/database"
	"go_mangasnake_api/api/handlers"
	"go_mangasnake_api/api/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	r := gin.Default()

	// rotas publicas
	r.POST("/token", handlers.GenerateAdminToken)
	adminRoutes := r.Group("/")
	adminRoutes.Use(middleware.AdminAuthMiddleware())
	adminRoutes.GET("/mangas", handlers.GetMangas)

	r.POST("/register", handlers.RegisterUser)
	r.POST("/login", handlers.LoginUser)
	// rotas protegidas
	protected := r.Group("/", middleware.JWTAuthMiddleware())
	{
		protected.POST("/manga", handlers.CreateManga)
		protected.GET("/manga/:id", handlers.GetManga)
		protected.PUT("/manga/:id", handlers.UpdateManga)
		protected.DELETE("/manga/:id", handlers.DeleteManga)
		protected.DELETE("/user/:id", handlers.DeleteUser)

	}

	r.Run(":8080")

}
