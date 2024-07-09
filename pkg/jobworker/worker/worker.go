package worker

import (
	"log"
	"sync"

	loggerpkg "github.com/DaZZler12/sjf-be/pkg/entities/logger"
	"github.com/DaZZler12/sjf-be/pkg/entities/sjf/model"
	sjfServicePkg "github.com/DaZZler12/sjf-be/pkg/entities/sjf/service"
	sjfJobQueuePkg "github.com/DaZZler12/sjf-be/pkg/jobworker/sjfjobqueue"
	"go.uber.org/zap"
)

type JobWorker interface {
	Start()
	Run()
	Stop()
	ProcessJob(job *model.SJF)          // ProcessJob processes a job
	AddJob(job *model.SJF)              // AddJob adds a job to the queue
	UpdateJobInDB(job *model.SJF) error // UpdateJobInDB updates a job in the database
	PrintQuqueStatus()                  // PrintCurrentQuque prints the current queue information

	// TODO: Implement these methods
	// RemoveJob(job *model.SJF) // RemoveJob removes a job from the queue
}

type Worker struct {
	JobsQueue  *sjfJobQueuePkg.QueueForSJF // JobsQueue is the queue of jobs
	lock       sync.Mutex                  // Lock
	stopChan   chan struct{}               // used to signal the worker to stop
	logger     *zap.Logger
	sjfService sjfServicePkg.Service // sjf-Service

}

var (
	JobWorkerInstance *Worker
	once              sync.Once
)

func NewWorker() *Worker {
	once.Do(func() {
		JobWorkerInstance = &Worker{
			JobsQueue:  &sjfJobQueuePkg.QueueForSJF{},
			lock:       sync.Mutex{},
			stopChan:   make(chan struct{}),
			logger:     loggerpkg.GetLoggerInstance(),
			sjfService: sjfServicePkg.New(),
		} // Initialize the worker
		JobWorkerInstance.JobsQueue.InitSJFQueue() // Initialize the queue
		if JobWorkerInstance.JobsQueue == nil {
			log.Fatal("JobsQueue is nil, failed to initialize the queue and start the worker")
		}
		if JobWorkerInstance.logger == nil {
			log.Fatal("Logger is nil, failed to initialize the logger and start the worker")
		} else {
			JobWorkerInstance.logger.Info("Job Worker initialized successfully")
		}

	})
	return JobWorkerInstance
}
