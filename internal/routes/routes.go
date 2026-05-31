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
}
