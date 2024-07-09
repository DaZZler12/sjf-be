package error

const (
	JobNotFound         string = "job not found"
	JobExists           string = "job already exists"
	JobNameEmpty        string = "job name cannot be empty"
	JobDurationEmpty    string = "job duration cannot be empty"
	JobDurationInvalid  string = "job duration should be greater than 0"
	JobCreateError      string = "error creating the job"
	JobUpdateError      string = "error updating the job"
	JobDeleteError      string = "error deleting the job"
	JobGetError         string = "error getting the job"
	JobListError        string = "error listing the jobs"
	JobInvalidID        string = "invalid job ID"
	ProcessError        string = "error processing the request"
	BindingError        string = "error binding the request"
	InternalServerError string = "internal server error"
	JobNameInvalid      string = "job name is invalid, it should have a-z, A-Z, 0-9, and _ only"
	ErrInvalidJob       string = "invalid job"
)
