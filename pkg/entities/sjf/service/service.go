package service

import (
	"context"
	"sync"

	"github.com/DaZZler12/sjf-be/pkg/entities/sjf/model"
	"github.com/DaZZler12/sjf-be/pkg/entities/sjf/store"
	"go.mongodb.org/mongo-driver/bson"
)

type Service interface {
	Create(ctx context.Context, sjf *model.SJF) (*model.SJF, error)
	// List(ctx context.Context) ([]*model.SJF, error)
	Get(ctx context.Context, filters *bson.M) (*model.SJF, error)
	// Update(ctx context.Context, sjf *model.SJF) error
	// Delete(ctx context.Context, id string) error
	CountDocuments(ctx context.Context, filter bson.M) (int64, error)
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
