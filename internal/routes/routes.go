package routes

import (
	"pg_pritani/backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	r *gin.Engine,
	authHandler *handler.AuthHandler,
) {
	r.POST("/auth/register", authHandler.Register)
	r.POST("/auth/login", authHandler.Login)
}
