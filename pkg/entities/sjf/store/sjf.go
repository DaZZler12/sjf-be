package store

import (
	"context"

	"github.com/DaZZler12/sjf-be/pkg/entities/sjf/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Create creates a new SJF process in the database
func (store *SJFStore) Create(ctx context.Context, sjf *model.SJF) (*model.SJF, error) {
	result, err := store.sjfCollection.InsertOne(ctx, sjf, &options.InsertOneOptions{})
	if err != nil {
		return nil, err
	}
	sjf.ID = result.InsertedID.(primitive.ObjectID)
	return sjf, nil
}

// List returns all the SJF processes in the database
func (store *SJFStore) List(ctx context.Context) ([]*model.SJF, error) {
	// todo: implement the list method
	return nil, nil
}

// Get returns the SJF process with the given filters
func (store *SJFStore) Get(ctx context.Context, filters *primitive.M) (*model.SJF, error) {
	obtainedSJF := &model.SJF{}
	err := store.sjfCollection.FindOne(ctx, obtainedSJF, filters, &options.FindOneOptions{})
	if err != nil {
		return nil, err
	}
	return obtainedSJF, nil

}

// Update updates the SJF process with the given ID
func (store *SJFStore) Update(ctx context.Context, sjf *model.SJF) error {
	// todo: implement the update method
	return nil
}

// Delete deletes the SJF process with the given ID
func (store *SJFStore) Delete(ctx context.Context, id string) error {
	// todo: implement the delete method
	return nil
}

// CountDocuments returns the number of documents in the SJF collection
func (store *SJFStore) CountDocuments(ctx context.Context, filter bson.M) (int64, error) {
	count, err := store.sjfCollection.CountDocuments(ctx, filter, &options.CountOptions{})
	if err != nil {
		return 0, err
	}
	return count, nil
}
