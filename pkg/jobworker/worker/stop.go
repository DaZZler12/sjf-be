package worker

func (w *Worker) Stop() {
	w.logger.Info("Stopping the job worker")
	close(w.stopChan) // Close the stop channel
}
