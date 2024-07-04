package service

import (
	"context"

	"github.com/DaZZler12/sjf-be/pkg/entities/sjf/model"
)

func (service *SJFService) Create(ctx context.Context, sjf *model.SJF) (*model.SJF, error) {
	return service.sjfStore.Create(ctx, sjf)
}
