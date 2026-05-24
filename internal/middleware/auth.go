package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authrization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "authorization header is rerquired",
			})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "invalid authorization format",
			})
			return
		}

		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "invalid or expired token",
			})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "invalid token claims",
			})
			return
		}

		ctx.Set("user_id", int(claims["user_id"].(float64)))
		ctx.Set("role", claims["role"].(string))
		ctx.Next()
	}
}
