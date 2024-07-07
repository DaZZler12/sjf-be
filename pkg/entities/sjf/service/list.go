package service

import (
	"context"

	"github.com/DaZZler12/sjf-be/pkg/entities/sjf/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (sjfService *SJFService) List(ctx context.Context, filters *bson.M, findOptions *options.FindOptions) ([]*model.SJF, error) {
	if filters == nil {
		filters = &bson.M{}
	}
	if findOptions == nil {
		findOptions = &options.FindOptions{}
	}
	return sjfService.sjfStore.List(ctx, filters, findOptions)
}
