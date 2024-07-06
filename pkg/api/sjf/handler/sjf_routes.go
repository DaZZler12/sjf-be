package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SjfInit(route *gin.RouterGroup) {
	// Initialize the handler
	sjfHandler := New()
	if sjfHandler == nil {
		log.Fatal("Failed to initialize the SJF handler, exiting...")
	}
	// define the routes
	sjfGroup := route.Group("/sjf")
	{
		version1 := sjfGroup.Group("/v1")
		{
			version1.POST("", sjfHandler.Create)
			// version1.GET("/list", sjfHandler.List)
			version1.GET("/:id", sjfHandler.Get)
			// version1.PUT("/update", sjfHandler.Update)
			version1.DELETE("/:id", sjfHandler.Delete)
		}
	}
	zap.S().Info("SJF routes initialized")
}
