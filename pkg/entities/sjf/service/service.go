package service

import (
	"context"
	"sync"

	"github.com/DaZZler12/sjf-be/pkg/entities/sjf/model"
	"github.com/DaZZler12/sjf-be/pkg/entities/sjf/store"
)

type Service interface {
	Create(ctx context.Context, sjf *model.SJF) (*model.SJF, error)
	// List(ctx context.Context) ([]*model.SJF, error)
	// GetByID(ctx context.Context, id string) (*model.SJF, error)
	// Update(ctx context.Context, sjf *model.SJF) error
	// Delete(ctx context.Context, id string) error
}

type SJFService struct {
	sjfStore store.Store
	// other resources can be added here, if needed
}

var (
	sjfService *SJFService
	once       sync.Once
)

// New creates a new service instance
func New() Service {
	once.Do(func() {
		sjfStore := store.New()
		if sjfStore == nil {
			return
		}
		sjfService = &SJFService{
			sjfStore: sjfStore,
		}
	})
	return sjfService
}