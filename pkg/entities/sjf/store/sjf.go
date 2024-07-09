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
func (store *SJFStore) List(ctx context.Context, filters *bson.M, findOptions *options.FindOptions) ([]*model.SJF, error) {
	obtainedSJFList := []*model.SJF{} // empty slice
	err := store.sjfCollection.FindMany(ctx, &obtainedSJFList, filters, findOptions)
	if err != nil {
		return nil, err
	}
	return obtainedSJFList, nil
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

// Update updates the SJF process with the given filters
func (store *SJFStore) Update(ctx context.Context, filters *bson.M, updates *bson.M) error {
	_, err := store.sjfCollection.UpdateOne(ctx, filters, updates, &options.UpdateOptions{})
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a SJF process with the given filters
func (store *SJFStore) Delete(ctx context.Context, filters *bson.M) error {
	_, err := store.sjfCollection.DeleteOne(ctx, filters, &options.DeleteOptions{})
	if err != nil {
		return err
	}
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
