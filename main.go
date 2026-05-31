package main

import (
	"os"
	"pg_pritani/backend/internal/db"
	"pg_pritani/backend/internal/handler"
	"pg_pritani/backend/internal/repository"
	"pg_pritani/backend/internal/routes"
	"pg_pritani/backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if err := godotenv.Load(); err != nil {
		log.Fatal().Msg("error loading .env")
	}

	database := db.ConnectDB()

	userRepo := repository.NewUserRepository(database)
	authService := service.NewAuthService(database, userRepo)
	authHandler := handler.NewAuthHandler(authService)

	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	adminRepo := repository.NewAdminRepository(database)
	adminService := service.NewAdminService(adminRepo)
	adminHandler := handler.NewAdminHandler(adminService)

	employeeRepo := repository.NewEmployeeRepository(database)
	employeeService := service.NewEmployeeService(employeeRepo)
	employeeHandler := handler.NewEmployeeHandler(employeeService)

	r := gin.Default()
	routes.SetupRoutes(r, authHandler, userHandler, adminHandler, employeeHandler)

	port := os.Getenv("APP_PORT")
	log.Info().Str("port", port).Msg("server running")

	if err := r.Run(":" + port); err != nil {
		log.Fatal().Err(err).Msg("error running server")
	}
}
