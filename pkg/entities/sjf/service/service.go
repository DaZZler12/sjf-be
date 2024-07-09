package service

import (
	"context"
	"sync"

	"github.com/DaZZler12/sjf-be/pkg/entities/sjf/model"
	"github.com/DaZZler12/sjf-be/pkg/entities/sjf/store"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service interface {
	Create(ctx context.Context, sjf *model.SJF) (*model.SJF, error)
	List(ctx context.Context, filters *bson.M, findOptions *options.FindOptions) ([]*model.SJF, error)
	Get(ctx context.Context, filters *bson.M) (*model.SJF, error)
	Update(ctx context.Context, sjf *model.SJF, filters *bson.M, updates *bson.M) error
	Delete(ctx context.Context, filters *bson.M) error
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
