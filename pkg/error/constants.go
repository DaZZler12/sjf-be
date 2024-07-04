package error

const (
	JobNotFound         = "job not found"
	JobExists           = "job already exists"
	JobNameEmpty        = "job name cannot be empty"
	JobDurationEmpty    = "job duration cannot be empty"
	JobDurationInvalid  = "job duration should be greater than 0"
	JobCreateError      = "error creating the job"
	JobUpdateError      = "error updating the job"
	JobDeleteError      = "error deleting the job"
	JobGetError         = "error getting the job"
	JobListError        = "error listing the jobs"
	JobInvalidID        = "invalid job ID"
	ProcessError        = "error processing the request"
	BindingError        = "error binding the request"
	InternalServerError = "internal server error"
)
