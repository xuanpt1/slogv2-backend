package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func CorsMiddleware() gin.HandlerFunc {
	return cors.New(
		cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
			AllowHeaders:     []string{"*", "Authorization"},
			ExposeHeaders:    []string{"Content-Length", "Content-Type", "text/plain", "Authorization"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		},
	)
}
