package worker

import (
	"container/heap"

	"github.com/DaZZler12/sjf-be/pkg/entities/sjf/model"
	"go.uber.org/zap"
)

func (w *Worker) AddJob(job *model.SJF) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.logger.Info("Adding a job to priority queue for processing names : ", zap.String("job name: ", job.Name))
	// w.JobsQueue.Push(job) // Push the job into the queue , this is not the correct way to push the job into the queue
	heap.Push(w.JobsQueue, job)
}
