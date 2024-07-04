package handler

import (
	"context"
	"net/http"

	requestModel "github.com/DaZZler12/sjf-be/pkg/api/sjf/model"
	"github.com/DaZZler12/sjf-be/pkg/entities/sjf/constants"
	"github.com/DaZZler12/sjf-be/pkg/entities/sjf/model"
	commonErrors "github.com/DaZZler12/sjf-be/pkg/error"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

/**
 * @api {post} /sjf Create a new SJF
 * @apiName CreateSJF
 * @apiGroup SJF
 * @apiVersion  0.1.0
 *
 * @apiParam  {String} name Name of the SJF process
 * @apiParam  {Number} duration Duration of the SJF process
 * @apiParamExample  {json} Request-Example:
 * {
 *     "name": "Process 1",
 *     "duration": 10
 * }
 *
 * @apiSuccess (201) {String} id ID of the SJF process
 * @apiSuccess (201) {String} name Name of the SJF process
 * @apiSuccess (201) {Number} duration Duration of the SJF process
 * @apiSuccess (201) {String} status Status of the SJF process
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 201 Created
 * {
 *     "id": "5f7b1b7b7f7b1b7b7f7b1b7b",
 *     "name": "Process 1",
 *     "duration": 10,
 *     "status": "pending"
 * }
 *
 * @apiError (400) {String} error Bad Request
 * @apiError (500) {String} error Internal Server Error
 * @apiErrorExample {json} Error-Response:
 * HTTP/1.1 400 Bad Request
 * {
 *     "error": "Bad Request"
 * }
 */

// Create creates a new SJF
// this is the handler function that will be called when the create endpoint is hit
func (handler *SJFHandler) Create(c *gin.Context) {
	ctx := context.Context(c)

	// from the context get the request body and Bind it to the SJF struct
	var sjfRequest *requestModel.SJFRequest
	if err := c.BindJSON(&sjfRequest); err != nil {
		handler.logger.Error("Failed to bind the request body", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": commonErrors.BindingError})
		return
	}

	// validate the request
	if err := sjfRequest.Validate(); err != nil {
		handler.logger.Error("Failed to validate the request", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sjfRequest.CalculateDuration()
	handler.logger.Info("Creating a new SJF", zap.Any("sjfRequest", sjfRequest))

	sjf := &model.SJF{
		ID:       primitive.NewObjectID(),
		Name:     sjfRequest.Name,
		Duration: sjfRequest.Duration,
		Status:   constants.Pending,
	}

	sjf, err := handler.sjfService.Create(ctx, sjf)
	if err != nil {
		handler.logger.Error("Failed to create the SJF", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": commonErrors.InternalServerError})
		return
	}
	handler.logger.Info("SJF created successfully", zap.Any("sjf", sjf))
	c.JSON(http.StatusCreated, sjf) // return the created SJF
}
