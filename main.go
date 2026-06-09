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

	productRepo := repository.NewProductRepository(database)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	supplierRepo := repository.NewSupplierRepository(database)
	supplierService := service.NewSupplierService(supplierRepo)
	supplierHandler := handler.NewSupplierHandler(supplierService)

	customerRepo := repository.NewCustomerRepository(database)
	customerService := service.NewCustomerService(customerRepo)
	customerHandler := handler.NewCustomerHandler(customerService)

	salesRepo := repository.NewSalesRepository(database)
	salesService := service.NewSalesService(database, salesRepo, productRepo)
	salesHandler := handler.NewSalesHandler(salesService)

	paymentRepo := repository.NewPaymentRepository(database)
	paymentService := service.NewPaymentService(database, paymentRepo, salesRepo)
	paymentHandler := handler.NewPaymentHandler(paymentService)

	r := gin.Default()
	routes.SetupRoutes(r, authHandler, userHandler, adminHandler, employeeHandler, productHandler, supplierHandler, customerHandler, salesHandler, paymentHandler)

	port := os.Getenv("APP_PORT")
	log.Info().Str("port", port).Msg("server running")

	if err := r.Run(":" + port); err != nil {
		log.Fatal().Err(err).Msg("error running server")
	}
}
