package service

import (
	"context"

	"github.com/DaZZler12/sjf-be/pkg/entities/sjf/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (service *SJFService) Get(ctx context.Context, filters *bson.M) (*model.SJF, error) {
	if filters == nil {
		filters = &bson.M{}
	}
	return service.sjfStore.Get(ctx, filters)
}
