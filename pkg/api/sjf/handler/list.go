package handler

import (
	"net/http"
	"strconv"

	commonErrors "github.com/DaZZler12/sjf-be/pkg/error"
	"github.com/DaZZler12/sjf-be/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

// List lists the SJF processes
// - search: this will be used to search the job name in the job name field of the job
// - _start: this will be used to start the pagination from the given index
// - _end: this will be used to end the pagination at the given index
// - sort: this will be used to sort the jobs based on the field name provided in the query param
// - order: this will be used to sort the jobs in ascending or descending order
// - status: this will be used to filter the jobs based on the status of the job
// @Summary List SJF processes
// @Description List SJF processes
// @Tags SJF
// @Accept json
// @Produce json
// @Param search query string false "Search by name"
// @Param _start query int false "Start index for pagination"
// @Param _end query int false "End index for pagination"
// @Param sort query string false "Sort by field"
// @Param order query string false "Sort order"
// @Param status query string false "Filter by status"
// @Success 200 {array} SJF
// @Router /api/sjf/v1 [get]

func (handler *SJFHandler) List(c *gin.Context) {

	// Default values
	defaultStart := int64(0)
	defaultEnd := int64(10)
	defaultSortField := "_id"
	defaultSortOrder := 1 // Ascending order

	// 1. filters
	filters := &bson.M{}
	// 2. findOptions
	findOptions := options.Find().SetSort(bson.D{{Key: defaultSortField, Value: defaultSortOrder}})

	// remember I need produciton level code..
	// search
	search := c.Query("search")
	if search != "" {
		safeString := utils.SanitizeString(search)
		handler.logger.Info("Sanitized search string", zap.String("search", safeString))
		(*filters)["name"] = bson.M{"$regex": safeString, "$options": "i"}
	}
	// status
	status := c.Query("status")
	if status != "" {
		(*filters)["status"] = status
	}

	// pagination
	startStr := c.Query("_start")
	endStr := c.Query("_end")
	start, err := strconv.ParseInt(startStr, 10, 64)
	if err != nil {
		start = defaultStart
	}
	end, err := strconv.ParseInt(endStr, 10, 64)
	if err != nil {
		end = defaultEnd
	}
	limit := end - start
	if limit > 0 {
		findOptions.SetSkip(start).SetLimit(limit)
	}

	// Sorting
	sort := c.Query("sort")
	if sort != "" {
		sortOrder := defaultSortOrder
		if c.Query("order") == "DESC" {
			sortOrder = -1
		}
		findOptions.SetSort(bson.D{{Key: sort, Value: sortOrder}})
	}

	totalCount, err := sjfHandler.sjfService.CountDocuments(c, *filters)
	if err != nil {
		handler.logger.Error("Failed to count the documents", zap.Error(err))
	}

	listOfJobs, err := sjfHandler.sjfService.List(c, filters, findOptions)
	if err != nil {
		handler.logger.Error("Failed to list the jobs", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": commonErrors.InternalServerError})
		return
	}
	c.Header("X-Total-Count", strconv.FormatInt(totalCount, 10))
	c.JSON(http.StatusOK, listOfJobs)
}
