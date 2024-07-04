package model

import (
	"errors"
	"fmt"
	"time"

	commonErrors "github.com/DaZZler12/sjf-be/pkg/error"
)

type SJFRequest struct {
	Name        string        `json:"name" validate:"required"`
	DurationSec float64       `json:"duration" validate:"required"`
	Duration    time.Duration // this will be calculated from DurationSec, not a part of request
}

func (sjfRequest *SJFRequest) CalculateDuration() {
	fmt.Println("Calculating Duration", time.Duration(sjfRequest.DurationSec)*time.Second)
	sjfRequest.Duration = time.Duration(sjfRequest.DurationSec * float64(time.Second))
}

func (sjfRequest *SJFRequest) Validate() error {
	if sjfRequest.Name == "" {
		return errors.New(commonErrors.JobNameEmpty)
	}
	if sjfRequest.DurationSec <= 0 {
		return errors.New(commonErrors.JobDurationInvalid)
	}
	return nil
}
