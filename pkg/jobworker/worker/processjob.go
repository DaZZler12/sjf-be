package worker

import (
	"time"

	"github.com/DaZZler12/sjf-be/pkg/entities/sjf/constants"
	"github.com/DaZZler12/sjf-be/pkg/entities/sjf/model"
)

func (w *Worker) ProcessJob(job *model.SJF) {
	defer func() {
		if r := recover(); r != nil {
			w.logger.Sugar().Errorf("Recovered in ProcessJob, %v", r)
		}
	}()

	w.logger.Sugar().Infof("Processing job %s", job.Name)
	job.Status = constants.Running
	w.UpdateJobInDB(job)
	time.Sleep(job.Duration) // Ensure this is a reasonable value
	job.Status = constants.Completed
	if err := w.UpdateJobInDB(job); err != nil {
		w.logger.Sugar().Errorf("Error updating job %s in the database: %v", job.Name, err)
		return
	}
	w.logger.Sugar().Infof("Job %s completed", job.Name)
}

// ref. function. First iteration

// func (w *Worker) ProcessJob(job *model.SJF) {
// 	w.logger.Infof("Processing job %s", job.Name)
// 	job.Status = constants.Running
// 	w.UpdateJobInDB(job)
// 	time.Sleep(job.Duration)
// 	job.Status = constants.Completed
// 	err := w.UpdateJobInDB(job)
// 	if err != nil {
// 		w.logger.Errorf("Error updating job %s in the database: %v", job.Name, err)
// 		return
// 	}
// 	w.logger.Infof("Job %s completed", job.Name)
// }
