package handler

import (
	"sync"

	loggerpkg "github.com/DaZZler12/sjf-be/pkg/entities/logger"
	"github.com/DaZZler12/sjf-be/pkg/entities/sjf/service"
	"github.com/DaZZler12/sjf-be/pkg/jobworker/worker"
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
	sjfWorker *worker.Worker
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
		sjfHandler = &SJFHandler{
			sjfService: sjfService,
			logger:     loggerpkg.GetLoggerInstance(),
			sjfWorker:  worker.NewWorker(),
		}
	})
	return sjfHandler
}
