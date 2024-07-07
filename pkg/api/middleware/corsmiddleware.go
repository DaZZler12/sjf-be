package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORSMiddleware is a middleware to handle CORS
//   - allows all origins, methods, and headers
//   - also allows credentials
//   - sets the max age to 12 hours
//   - returns a gin.HandlerFunc
func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-BrowserFingerprint", "X-Workspace-ID"},
		ExposeHeaders:    []string{"Content-Length", "Content-Range", "X-Total-Count"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
