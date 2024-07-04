package store

import (
	"context"

	"github.com/DaZZler12/sjf-be/pkg/entities/sjf/model"
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

// GetByID returns the SJF process with the given ID
func (store *SJFStore) GetByID(ctx context.Context, id string) (*model.SJF, error) {
	// todo: implement the get by id method
	return nil, nil
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
