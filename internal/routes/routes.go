package routes

import (
	"pg_pritani/backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	r *gin.Engine,
	authHandler *handler.AuthHandler,
	userHandler *handler.UserHandler,
	adminHandler *handler.AdminHandler,
	employeeHandler *handler.EmployeeHandler,
	productHandler *handler.ProductHandler,
	supplierHandler *handler.SupplierHandler,
	customerHandler *handler.CustomerHandler,
	salesHandler *handler.SalesHandler,
) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/auth/login", authHandler.Login)
	}
	user := r.Group("/user")
	{
		user.GET("", userHandler.GetAll)
		user.GET("/:id", userHandler.GetByID)
		user.PATCH("/:id", userHandler.Update)
		user.DELETE("/:id", userHandler.Delete)
	}

	admin := r.Group("/admin")
	{
		admin.GET("", adminHandler.GetAll)
		admin.GET("/:id", adminHandler.GetByID)
		admin.PATCH("/:id", adminHandler.Update)
		admin.DELETE("/:id", adminHandler.Delete)
	}

	employee := r.Group("/employee")
	{
		employee.GET("", employeeHandler.GetAll)
		employee.GET("/:id", employeeHandler.GetByID)
		employee.PATCH("/:id", employeeHandler.Update)
		employee.DELETE("/:id", employeeHandler.Delete)
	}

	product := r.Group("/products")
	{
		product.GET("", productHandler.GetAll)
		product.GET("/:id", productHandler.GetByID)
		product.POST("", productHandler.Create)
		product.PATCH("/:id", productHandler.Update)
		product.DELETE("/:id", productHandler.Delete)
	}

	supplier := r.Group("/suppliers")
	{
		supplier.GET("", supplierHandler.GetAll)
		supplier.GET("/:id", supplierHandler.GetByID)
		supplier.POST("", supplierHandler.Create)
		supplier.PATCH("/:id", supplierHandler.Update)
		supplier.DELETE("/:id", supplierHandler.Delete)
	}

	customer := r.Group("/customers")
	{
		customer.GET("", customerHandler.GetAll)
		customer.GET("/:id", customerHandler.GetByID)
		customer.POST("", customerHandler.Create)
		customer.PATCH("/:id", customerHandler.Update)
		customer.DELETE("/:id", customerHandler.Delete)

	}

	sales := r.Group("/sales")
	{
		sales.GET("", salesHandler.GetAll)
		sales.GET("/:id", salesHandler.GetByID)
		sales.POST("", salesHandler.Create)
		sales.POST("/:id/payment", salesHandler.AddPayment)
		sales.DELETE("/:id", salesHandler.Delete)

	}
}
