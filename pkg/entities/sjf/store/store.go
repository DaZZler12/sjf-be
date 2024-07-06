package store

// here I will define the store interface and also
// we need to define the struct that will implement this interface
// this will be used to interact with the database

import (
	"context"
	"log"
	"sync"

	"github.com/DaZZler12/sjf-be/pkg/config"
	"github.com/DaZZler12/sjf-be/pkg/entities/database/mongo"
	mongoORM "github.com/DaZZler12/sjf-be/pkg/entities/database/mongo/orm"
	"github.com/DaZZler12/sjf-be/pkg/entities/schema"
	"github.com/DaZZler12/sjf-be/pkg/entities/sjf/model"
	"go.mongodb.org/mongo-driver/bson"
)

// Store is an interface that defines the methods that should be implemented by the store struct
type Store interface {
	Create(ctx context.Context, sjf *model.SJF) (*model.SJF, error)
	List(ctx context.Context) ([]*model.SJF, error)
	Get(ctx context.Context, filters *bson.M) (*model.SJF, error)
	Update(ctx context.Context, sjf *model.SJF) error
	Delete(ctx context.Context, filters *bson.M) error
	CountDocuments(ctx context.Context, filter bson.M) (int64, error)
}

// store is a struct that will implement the Store interface
type SJFStore struct {
	// Here we can define the database connection or any other fields that are required
	sjfCollection mongoORM.MongoCollectionDerived
}

var (
	store *SJFStore
	once  sync.Once
)

// New creates a new store instance
func New() Store {
	once.Do(func() {
		config, _ := config.LoadConfig("config/")
		_, _ = mongo.GetMongoDBInstance(&config.Database)
		collection := mongo.GetMongoCollection(schema.SJFCollection)
		if collection == nil {
			log.Fatalf("Failed to get the SJF collection")
		}
		// Create a unique index on the 'name' field
		indexModel := bson.D{{Key: "name", Value: 1}}
		if err := mongo.CreateUniqueIndexOnAFields(collection, indexModel); err != nil {
			log.Fatalf("Failed to create a unique index on the 'name' field: %v", err)
		}

		store = &SJFStore{
			sjfCollection: collection,
		}
	})
	return store

}
