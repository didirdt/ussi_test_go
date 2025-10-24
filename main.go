package main

import (
	"fmt"
	"log"
	"ussi_test/config"
	"ussi_test/controllers"
	"ussi_test/database"
	"ussi_test/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database
	if err := database.InitDatabase(); err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	defer database.CloseDatabase()

	// Initialize JWT
	middleware.InitJWT(cfg)

	// Set Gin mode
	if gin.Mode() == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	ussiGo := gin.Default()
	ussiGo.Use(gin.Logger())
	ussiGo.Use(gin.Recovery())

	// Routes
	ussiGo.POST("/api/login", controllers.Login)
	ussiGo.POST("/api/register", controllers.Register)

	auth := ussiGo.Group("/api")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/profile", controllers.GetProfile)

		auth.GET("/notes", controllers.GetNotes)
		auth.POST("/notes", controllers.CreateNote)
		auth.GET("/note/:id", controllers.GetNote)
		auth.PUT("/note/:id", controllers.UpdateNote)
		auth.DELETE("/note/:id", controllers.DeleteNote)

		admin := auth.Group("/")
		admin.Use(middleware.AdminOnly())
		{
			admin.GET("/users", controllers.GetUsers)
			admin.PUT("/users/:id", controllers.UpdateUser)
			admin.DELETE("/users/:id", controllers.DeleteUser)
		}
	}

	// Start server
	serverAddr := fmt.Sprintf(":%s", cfg.Server.Port)
	fmt.Printf("Server starting on %s\n", serverAddr)
	if err := ussiGo.Run(serverAddr); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
