package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/DaZZler12/sjf-be/pkg/entities/sjf/model"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

// GetWebSocketData gets the WebSocket data for the given job IDs
// @Summary Get WebSocket data for the given job IDs
// @Description Get WebSocket data for the given job IDs
// @Tags SJF
// @Accept json
// @Produce json
// @Param jobIDs body []string true "Job IDs"
// @Success 200 {object} SJF
// @Router /api/sjf/v1/websocket [get]

func (handler *SJFHandler) GetWebSocketDataForJobIDs(c *gin.Context) {
	conn, err := handler.websocketUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		handler.logger.Error("Failed to set WebSocket upgrade", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade WebSocket"})
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			handler.logger.Error("Failed to read WebSocket message", zap.Error(err))
			break
		}

		jobIDs := []string{}
		if err := json.Unmarshal(message, &jobIDs); err != nil {
			handler.logger.Error("Failed to unmarshal job IDs", zap.Error(err))
			conn.WriteJSON(gin.H{"error": "Failed to unmarshal job IDs"})
			continue
		}
		for _, id := range jobIDs {
			go handler.sendJobStatus(conn, id) // goroutine for each job ID, this will help in parallel processing
		}
	}
}

func (handler *SJFHandler) sendJobStatus(conn *websocket.Conn, id string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sjf, err := handler.fetchJobStatus(ctx, id)
	if err != nil {
		handler.logger.Error("Failed to retrieve SJF process", zap.Error(err))
		conn.WriteJSON(gin.H{"error": "Failed to retrieve job status", "id": id})
		return
	}

	if err = conn.WriteJSON(sjf); err != nil {
		handler.logger.Error("Failed to write WebSocket message", zap.Error(err))
	}
}

func (handler *SJFHandler) fetchJobStatus(ctx context.Context, id string) (*model.SJF, error) {
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filters := &bson.M{"_id": idObj}
	sjf, err := handler.sjfService.Get(ctx, filters)
	if err != nil {
		return nil, err
	}

	return sjf, nil
}
