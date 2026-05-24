package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role, exists := ctx.Get("role")
		if !exists {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "forbidden",
			})
			return
		}

		for _, r := range roles {
			if r == role.(string) {
				ctx.Next()
				return
			}
		}
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "you don't have permission to access this resource",
		})
	}

}
