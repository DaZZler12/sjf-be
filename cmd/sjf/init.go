package sjf

import (
	"log"

	"github.com/DaZZler12/sjf-be/pkg/api/routes"
	"github.com/DaZZler12/sjf-be/pkg/config"
	"github.com/DaZZler12/sjf-be/pkg/entities/database/mongo"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	ServerConfig *config.ServerConfig
)

func Init() (*gin.Engine, error) {
	configPath := "config/"
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		return nil, err
	}
	ServerConfig = &cfg.ServerConfig
	_, err = mongo.GetMongoDBInstance(&cfg.Database) // initialize the mongo db
	if err != nil {
		return nil, err
	}

	ginEngineInstance := routes.GenerateNewGinEngine(zap.L()) // this funciton is resposible for generating the gin engine and initializing the routes
	return ginEngineInstance, nil
}

func StartServer() {
	// get the server engine and start the server
	ginEngineInstance, err := Init()
	if err != nil {
		log.Fatal("Error initializing server", err) // failed to initialize the resource
	}
	zap.S().Info("Starting server on " + ServerConfig.Host + ":" + ServerConfig.Port)
	ginEngineInstance.Run(ServerConfig.Host + ":" + ServerConfig.Port)
}
