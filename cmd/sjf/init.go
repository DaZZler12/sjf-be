package sjf

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DaZZler12/sjf-be/pkg/api/routes"
	"github.com/DaZZler12/sjf-be/pkg/config"
	"github.com/DaZZler12/sjf-be/pkg/entities/database/mongo"
	loggerpkg "github.com/DaZZler12/sjf-be/pkg/entities/logger"
	"github.com/DaZZler12/sjf-be/pkg/jobworker/worker"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	ServerConfig       *config.ServerConfig
	jobsWorkerInstance *worker.Worker
	logger             *zap.Logger
)

func Init() (*gin.Engine, error) {
	configPath := "config/"
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		return nil, err
	}

	// configure the logger
	logger = loggerpkg.GetLoggerInstance()
	if logger == nil {
		log.Fatal("Failed to initialize logger")
	}
	defer logger.Sync() // flushes buffer, if any

	ServerConfig = &cfg.ServerConfig
	_, err = mongo.GetMongoDBInstance(&cfg.Database) // initialize the mongo db
	if err != nil {
		return nil, err
	}

	// initialize the job worker and start the worker
	jobsWorkerInstance = worker.NewWorker()
	if jobsWorkerInstance == nil {
		log.Fatal("Failed to initialize the job worker")
	}
	// start the worker
	jobsWorkerInstance.Start()

	ginEngineInstance := routes.GenerateNewGinEngine(logger)
	// this funciton is resposible for generating the gin engine and initializing the routes
	return ginEngineInstance, nil
}

func StartServer() {
	// get the server engine and start the server
	ginEngineInstance, err := Init()
	if err != nil {
		log.Fatal("Error initializing server", err)
	}

	// Setting up the server
	server := &http.Server{
		Addr:    ServerConfig.Host + ":" + ServerConfig.Port,
		Handler: ginEngineInstance, // Use ginEngineInstance as the handler
	}

	// Setting up a channel to listen for interrupt or termination signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Running the server in a goroutine so that it doesn't block
	go func() {
		logger.Debug("Starting server on " + server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server failed to start:", err)
		}
	}()

	// Block until a signal is received
	<-quit

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown the server gracefully.
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown Failed:", err)
	}

	// Stop the worker
	jobsWorkerInstance.Stop()
	logger.Info("Shutting down gracefully, worker stopped")
}

// TODO: need to use this when 2 differnt micorservices is being built for the same project

// func StartServer() {
// 	// get the server engine and start the server
// 	ginEngineInstance, err := Init()
// 	if err != nil {
// 		log.Fatal("Error initializing server", err) // failed to initialize the resource
// 	}
// 	zap.S().Info("Starting server on " + ServerConfig.Host + ":" + ServerConfig.Port)
// 	ginEngineInstance.Run(ServerConfig.Host + ":" + ServerConfig.Port)
// }
