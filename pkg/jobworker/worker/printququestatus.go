package worker

import (
	"github.com/DaZZler12/sjf-be/pkg/entities/sjf/model"
)

func (w *Worker) PrintQueueStatus() {
	defer func() {
		if r := recover(); r != nil {
			w.logger.Sugar().Errorf("Recovered in PrintQueueStatus, %v", r)
		}
	}()

	w.lock.Lock()
	defer w.lock.Unlock()
	w.logger.Info("Current Queue Status:")
	for i := 0; i < w.JobsQueue.Len(); i++ {
		job := w.JobsQueue.Get(i).(*model.SJF)
		if job == nil {
			w.logger.Sugar().Warnf("Job @ Position %d is nil", i+1)
			continue
		}
		w.logger.Sugar().Infof("\nJob @ Position %d: %s, Duration: %v, Priority: %d", i+1, job.Name, job.Duration)
	}
}
