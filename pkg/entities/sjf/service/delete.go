package service

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (sjfService *SJFService) Delete(ctx context.Context, filters *bson.M) error {
	if filters == nil {
		filters = &bson.M{}
	}
	_, err := sjfService.Get(ctx, filters)
	if err != nil {
		return err
	}
	return sjfService.sjfStore.Delete(ctx, filters)
}
