package constants

type ProcessStatus string

const (
	Pending   ProcessStatus = "pending"
	Running   ProcessStatus = "running"
	Completed ProcessStatus = "completed"
)
