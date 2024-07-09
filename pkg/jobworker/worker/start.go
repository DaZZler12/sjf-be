package worker

func (w *Worker) Start() {
	w.logger.Info("Starting the job worker")
	go w.Run()
}
