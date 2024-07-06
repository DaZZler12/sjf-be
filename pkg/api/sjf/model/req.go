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
	// the name should be have a-z, A-Z, 0-9, and _ only
	// if the name has any other characters, return an error
	if !isAlphaNumeric(sjfRequest.Name) {
		return errors.New(commonErrors.JobNameInvalid)
	}
	if sjfRequest.DurationSec <= 0 {
		return errors.New(commonErrors.JobDurationInvalid)
	}
	return nil
}

func isAlphaNumeric(s string) bool {
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			continue
		}
		if r >= 'A' && r <= 'Z' {
			continue
		}
		if r >= '0' && r <= '9' {
			continue
		}
		if r == '_' {
			continue
		}
		return false
	}
	return true
}
