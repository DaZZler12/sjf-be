package model

import (
	"time"

	"github.com/DaZZler12/sjf-be/pkg/entities/sjf/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SJF struct {
	ID       primitive.ObjectID      `json:"id" bson:"_id"`            // SJF process ID
	Name     string                  `json:"name" bson:"name"`         // SJF process name
	Duration time.Duration           `json:"duration" bson:"duration"` // SJF process duration
	Status   constants.ProcessStatus `json:"status" bson:"status"`     // SJF process status
}
