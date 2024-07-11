package handler

import (
	"net/http"
	"sync"

	loggerpkg "github.com/DaZZler12/sjf-be/pkg/entities/logger"
	"github.com/DaZZler12/sjf-be/pkg/entities/sjf/service"
	"github.com/DaZZler12/sjf-be/pkg/jobworker/worker"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type Handler interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Get(c *gin.Context)
	// Update(c *gin.Context)
	Delete(c *gin.Context)
	GetWebSocketDataForJobIDs(c *gin.Context)
	GetJobStatusUsingWebSocket(c *gin.Context)
}

type SJFHandler struct {
	sjfService service.Service
	logger     *zap.Logger
	// other resources can be added here, if needed
	sjfWorker         *worker.Worker      // Worker to process the jobs
	websocketUpgrader *websocket.Upgrader // Upgrader for WebSocket protocol
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
			sjfService:        sjfService,
			logger:            loggerpkg.GetLoggerInstance(),
			sjfWorker:         worker.NewWorker(),
			websocketUpgrader: WebSocketUpgrader(),
		}
	})
	return sjfHandler
}

// WebSocketUpgrader returns a new instance of websocket.Upgrader
// used to upgrade a GET request to WebSocket protocol
func WebSocketUpgrader() *websocket.Upgrader {
	return &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true // CORS policy can be adjusted here as needed
		},
	}
}
