package orm

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type MongoCollection interface {
	FindOne(ctx context.Context, filter interface{},
		opts ...*options.FindOneOptions) *mongo.SingleResult
	CountDocuments(ctx context.Context, filter interface{},
		opts ...*options.CountOptions) (int64, error)
	Find(ctx context.Context, filter interface{},
		opts ...*options.FindOptions) (*mongo.Cursor, error)
	InsertOne(ctx context.Context, document interface{},
		opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	InsertMany(ctx context.Context, documents []interface{},
		opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{},
		opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	DeleteOne(ctx context.Context, filter interface{},
		opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	Aggregate(ctx context.Context, filter interface{},
		opts ...*options.AggregateOptions) (*mongo.Cursor, error)
	Distinct(ctx context.Context, fieldName string, filter interface{},
		opts ...*options.DistinctOptions) ([]interface{}, error)
	BulkWrite(ctx context.Context, models []mongo.WriteModel,
		opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error)
	ReplaceOne(ctx context.Context, filter interface{},
		replacement interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error)
	DeleteMany(ctx context.Context, filter interface{},
		opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	Indexes() mongo.IndexView
	UpdateMany(ctx context.Context, filter interface{}, update interface{},
		opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
}

type MongoCollectionDerived interface {
	FindOne(ctx context.Context, target interface{}, filter interface{}, opts ...*options.FindOneOptions) error
	FindMany(ctx context.Context, target interface{}, filter interface{}, opts ...*options.FindOptions) error
	InsertOne(ctx context.Context, target interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	InsertMany(ctx context.Context, target []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	Aggregate(ctx context.Context, target interface{}, filter interface{}, opts ...*options.AggregateOptions) error
	Distinct(ctx context.Context, fieldName string, filter interface{}, opts ...*options.DistinctOptions) ([]interface{}, error)
	BulkWrite(ctx context.Context, models []mongo.WriteModel, opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error)
	CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error)
	ReplaceOne(ctx context.Context, filter interface{}, replacement interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error)
	DeleteMany(ctx context.Context, filter interface{},
		opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	Indexes() mongo.IndexView
	UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
}

type MongoORM struct {
	Col     MongoCollection
	ColName string
	Logger  *zap.Logger
}
