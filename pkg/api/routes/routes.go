package routes

import (
	"sync"
	"time"

	"github.com/DaZZler12/sjf-be/pkg/api/healthcheck"
	middleware "github.com/DaZZler12/sjf-be/pkg/api/middleware"
	sjfHandler "github.com/DaZZler12/sjf-be/pkg/api/sjf/handler"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	ginEngineInstance *gin.Engine
	once              sync.Once
)

// setupRoutes sets up the routes for the server
// this function is responsible for setting up the routes
// it will receive the gin engine and will setup the routes
func setupRoutes() {

	route := ginEngineInstance.Group("/api")
	route.GET("/", healthcheck.HealthCheck)
	// call the InitRoutes of each module to initialize the individual apis' routes
	sjfHandler.SjfInit(route)
}

// GenerateNewGinEngine
//   - it will return the gin engine instance
//   - this function is responsible for creating a new gin engine instance
//   - it will create a new gin engine instance and setup the routes
//   - it will also add the ginzap middleware to log the requests and panics
//   - it will also add the CORS middleware
//   - it will return the gin engine instance
func GenerateNewGinEngine(logger *zap.Logger) *gin.Engine {
	once.Do(func() {
		ginEngineInstance = gin.New() // create a new gin engine instance
		// Add a ginzap middleware, which:
		//   - Logs all requests, like a combined access and error log.
		//   - Logs to stdout.
		//   - RFC3339 with UTC time format.
		ginEngineInstance.Use(ginzap.Ginzap(logger, time.RFC3339, true))
		// Logs all panic to error log
		//   - stack means whether output the stack info.
		ginEngineInstance.Use(ginzap.RecoveryWithZap(logger, true))
		ginEngineInstance.Use(middleware.CORSMiddleware()) // as of now only CORS middleware is added
		setupRoutes()
	})
	return ginEngineInstance
}
