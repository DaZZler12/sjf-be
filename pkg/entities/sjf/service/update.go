package service

import (
	"context"

	"github.com/DaZZler12/sjf-be/pkg/entities/sjf/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (sjfService *SJFService) Update(ctx context.Context, sjf *model.SJF, filters *bson.M, updates *bson.M) error {
	if filters == nil {
		if sjf == nil {
			filters = &bson.M{} // default filter
		} else {
			filters = &bson.M{
				"_id": sjf.ID, // default filter by ID
			}
		}
	}
	return sjfService.sjfStore.Update(ctx, filters, updates)
}
