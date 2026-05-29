package main

import (
	"log"
	"os"
	"pg_pritani/backend/internal/db"
	"pg_pritani/backend/internal/handler"
	"pg_pritani/backend/internal/repository"
	"pg_pritani/backend/internal/routes"
	"pg_pritani/backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env")
	}
	database := db.ConnectDB()

	// repo
	userRepo := repository.NewUserRepository(database)
	authService := service.NewAuthService(database, userRepo)
	authHandler := handler.NewAuthHandler(authService)

	r := gin.Default()
	routes.SetupRoutes(r, authHandler)

	port := os.Getenv("APP_PORT")
	log.Printf("server running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
