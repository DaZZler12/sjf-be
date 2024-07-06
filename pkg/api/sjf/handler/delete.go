package handler

import (
	"net/http"

	commonerrors "github.com/DaZZler12/sjf-be/pkg/error"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (handler *SJFHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		handler.logger.Error("Invalid ID", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": commonerrors.JobInvalidID})
		return
	}
	filters := &bson.M{
		"_id": idObj,
	}

	err = handler.sjfService.Delete(c, filters)
	if err != nil {
		handler.logger.Error("Failed to delete SJF process", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": commonerrors.JobDeleteError})
		return
	}
	handler.logger.Info("SJF process deleted successfully", zap.String("id", id))
	c.JSON(http.StatusOK, gin.H{"message": "SJF process deleted successfully"})
}
