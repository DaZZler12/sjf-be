package handler

import (
	"time"

	"github.com/DaZZler12/sjf-be/pkg/entities/sjf/constants"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (handler *SJFHandler) GetJobStatusUsingWebSocket(c *gin.Context) {
	ctx := c.Request.Context()

	conn, err := handler.websocketUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		handler.logger.Error("Failed to upgrade to WebSocket", zap.Error(err))
		return
	}
	defer conn.Close()

	jobID := c.Param("id")
	if jobID == "" {
		handler.logger.Error("Job ID not provided")
		conn.WriteJSON(gin.H{"error": "Job ID not provided", "message": "Please provide a valid job ID"})
		return
	}

	jobObjID, err := primitive.ObjectIDFromHex(jobID)
	if err != nil {
		handler.logger.Error("Invalid job ID", zap.Error(err))
		conn.WriteJSON(gin.H{"error": "Invalid job ID", "message": "Please provide a valid job ID"})
		return
	}

	ticker := time.NewTicker(2 * time.Second) // Send job status every 2 seconds
	defer ticker.Stop()                       // Stop the ticker when the function returns

	for {
		select {
		case <-ticker.C:
			sjf, err := handler.sjfService.Get(ctx, &bson.M{"_id": jobObjID})
			if err != nil {
				handler.logger.Error("Failed to retrieve SJF process", zap.Error(err))
				conn.WriteJSON(gin.H{"error": "Failed to retrieve SJF process", "message": "Failed to retrieve SJF process"})
				return
			}
			if sjf.Status == constants.Completed {
				if err := conn.WriteJSON(gin.H{"jobinfo": sjf}); err != nil {
					handler.logger.Error("Failed to write message to WebSocket", zap.Error(err))
				}
				handler.logger.Info("Job completed, closing WebSocket connection")
				return // Close the connection as the job is completed
			}

			if err := conn.WriteJSON(gin.H{"jobinfo": sjf}); err != nil {
				handler.logger.Error("Failed to write message to WebSocket", zap.Error(err))
				return
			}
		case <-ctx.Done():
			return // Return if the context is cancelled
		}
	}
}
