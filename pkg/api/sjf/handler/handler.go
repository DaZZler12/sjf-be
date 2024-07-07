package handler

import (
	"sync"

	"github.com/DaZZler12/sjf-be/pkg/entities/sjf/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Get(c *gin.Context)
	// Update(c *gin.Context)
	Delete(c *gin.Context)
}

type SJFHandler struct {
	sjfService service.Service
	logger     *zap.Logger
	// other resources can be added here, if needed

}

var (
	sjfHandler *SJFHandler
	once       sync.Once
)

func New() Handler {
	once.Do(func() {
		sjfService := service.New()
		if sjfService == nil {
			return
		}
		productionLogger, err := zap.NewProduction()
		if err != nil {
			return
		}
		sjfHandler = &SJFHandler{
			sjfService: sjfService,
			logger:     productionLogger,
		}
	})
	return sjfHandler
}
