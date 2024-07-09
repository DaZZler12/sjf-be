package worker

import (
	"container/heap"
	"time"

	"github.com/DaZZler12/sjf-be/pkg/entities/sjf/model"
)

func (w *Worker) Run() {

	for {
		select {
		case <-w.stopChan:
			w.logger.Info("Stopping the job worker")
			return // Exit the loop and stop the worker
		default:
			w.lock.Lock()              // Lock the worker
			if w.JobsQueue.Len() > 0 { // Check if there are jobs in the queue
				// w.PrintQueueStatus()                  // Print the queue status // only for debugging can be removed
				job := heap.Pop(w.JobsQueue).(*model.SJF) // Pop a job from the queue, we will always have the shortest job first, as we are using the Priority Queue based on the duration
				w.lock.Unlock()
				w.ProcessJob(job)
			} else {
				w.logger.Info("No jobs in the queue, waiting for jobs to be added")
				w.lock.Unlock()             // Ensure the lock is released if no jobs are in the queue
				time.Sleep(2 * time.Second) // Sleep for a short time before checking the queue again
			}
		}
	}
}
