package service

import (
	"context"
	"fmt"

	"github.com/DaZZler12/sjf-be/pkg/entities/sjf/constants"
	"go.mongodb.org/mongo-driver/bson"
)

func (sjfService *SJFService) Delete(ctx context.Context, filters *bson.M) error {
	if filters == nil {
		filters = &bson.M{}
	}
	obtainedSJFJob, err := sjfService.Get(ctx, filters)
	if err != nil {
		return err
	}
	if obtainedSJFJob.Status == constants.Running || obtainedSJFJob.Status == constants.Pending {
		return fmt.Errorf("cannot delete a job that is in %s state", obtainedSJFJob.Status)
	}
	return sjfService.sjfStore.Delete(ctx, filters)
}
