package service

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (service *SJFService) CountDocuments(ctx context.Context, filter bson.M) (int64, error) {
	return service.sjfStore.CountDocuments(ctx, filter)
}
