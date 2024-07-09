package worker

import (
	"context"
	"errors"

	"github.com/DaZZler12/sjf-be/pkg/entities/sjf/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (w *Worker) UpdateJobInDB(job *model.SJF) error {
	w.lock.Lock()
	defer w.lock.Unlock()
	// Update the job in the database
	if job == nil {
		w.logger.Warn("Job is nil, unable to update job in the database")
		return errors.New("job is nil")
	}

	filters := &bson.M{
		"_id": job.ID, // default filter by ID
	}
	updates := &bson.M{
		"$set": &bson.M{
			"status": job.Status,
		},
	}
	err := w.sjfService.Update(context.Background(), job, filters, updates)
	if err != nil {
		w.logger.Sugar().Errorf("Error updating job in the database: %v", err)
		return err
	}
	w.logger.Sugar().Infof("Job %s updated in the database", job.Name)
	return nil
}
