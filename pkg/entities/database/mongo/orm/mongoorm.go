// implements the MongoCollectionDerived interface which is a derived interface from the MongoCollection interface.
// The MongoCollection interface defines the methods that can be used to interact with the MongoDB collection.
// The MongoCollectionDerived interface defines the methods that can be used to interact with the MongoDB collection in a more user-friendly way.
// The MongoORM struct is a struct that contains a MongoCollection interface and the name of the collection.
// The MongoORM struct is used to interact with the MongoDB collection in a more user-friendly way.
package orm

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var _ MongoCollectionDerived = &MongoORM{}

func (mongoOrm *MongoORM) FindOne(ctx context.Context, target interface{}, filter interface{}, opts ...*options.FindOneOptions) error {
	mongoOrm.Logger.Debug("findOneQuery", zap.Any("query", filter), zap.Any("opts", opts))

	result := mongoOrm.Col.FindOne(ctx, filter, opts...)
	if result.Err() != nil {
		return result.Err()
	}
	return result.Decode(target)
}

func (mongoOrm *MongoORM) FindMany(ctx context.Context, target interface{}, filter interface{}, opts ...*options.FindOptions) error {
	mongoOrm.Logger.Debug("findQuery", zap.Any("query", filter), zap.Any("opts", opts))

	cur, err := mongoOrm.Col.Find(ctx, filter, opts...)
	if err != nil {
		return err
	}
	return cur.All(ctx, target)
}

func (mongoOrm *MongoORM) InsertOne(ctx context.Context, target interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	mongoOrm.Logger.Debug("insertOne", zap.Any("target", target), zap.Any("opts", opts))
	return mongoOrm.Col.InsertOne(ctx, target, opts...)
}

func (mongoOrm *MongoORM) InsertMany(ctx context.Context, items []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	mongoOrm.Logger.Debug("insertMany", zap.Any("items", items), zap.Any("opts", opts))
	return mongoOrm.Col.InsertMany(ctx, items, opts...)
}

func (mongoOrm *MongoORM) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	mongoOrm.Logger.Debug("updateOne", zap.Any("filter", filter), zap.Any("update", update), zap.Any("opts", opts))
	return mongoOrm.Col.UpdateOne(ctx, filter, update, opts...)
}

func (mongoOrm *MongoORM) Aggregate(ctx context.Context, target interface{}, filter interface{}, opts ...*options.AggregateOptions) error {
	mongoOrm.Logger.Debug("aggregate filter", zap.Any("aggregate", filter))
	cur, err := mongoOrm.Col.Aggregate(ctx, filter, opts...)
	if err != nil {
		mongoOrm.Logger.Error("failed to run aggregate filter", zap.Error(err))
		return err
	}

	err = cur.All(ctx, target)
	if err != nil {
		mongoOrm.Logger.Error("failed to decode all result", zap.Error(err))
		return err
	}

	return nil

}

func (mongoOrm *MongoORM) Distinct(
	ctx context.Context,
	fieldName string,
	filter interface{},
	opts ...*options.DistinctOptions,
) ([]interface{}, error) {
	mongoOrm.Logger.Debug("distinct query", zap.Any("query_query", filter))
	return mongoOrm.Col.Distinct(ctx, fieldName, filter, opts...)
}

func (mongoOrm *MongoORM) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	mongoOrm.Logger.Debug("delete query", zap.Any("delete_query", filter))
	return mongoOrm.Col.DeleteOne(ctx, filter, opts...)
}

func (mongoOrm *MongoORM) BulkWrite(
	ctx context.Context, models []mongo.WriteModel, opts ...*options.BulkWriteOptions,
) (*mongo.BulkWriteResult, error) {
	mongoOrm.Logger.Debug("bulk write", zap.Any("models", models))
	return mongoOrm.Col.BulkWrite(ctx, models, opts...)
}

func (mongoOrm *MongoORM) CountDocuments(
	ctx context.Context, filter interface{}, opts ...*options.CountOptions,
) (int64, error) {
	mongoOrm.Logger.Debug("count documents", zap.Any("filter", filter))
	return mongoOrm.Col.CountDocuments(ctx, filter, opts...)
}

func (mongoOrm *MongoORM) ReplaceOne(ctx context.Context, filter interface{}, replacement interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
	mongoOrm.Logger.Debug("replace one", zap.Any("filter", filter), zap.Any("replacement", replacement))
	return mongoOrm.Col.ReplaceOne(ctx, filter, replacement, opts...)
}

func (mongoOrm *MongoORM) DeleteMany(ctx context.Context, filter interface{},
	opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	mongoOrm.Logger.Debug("delete many", zap.Any("filter", filter))
	return mongoOrm.Col.DeleteMany(ctx, filter, opts...)

}
func (mongoOrm *MongoORM) Indexes() mongo.IndexView {
	mongoOrm.Logger.Debug("indexes")
	return mongoOrm.Col.Indexes()
}

func (mongoOrm *MongoORM) UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return mongoOrm.Col.UpdateMany(ctx, filter, update, opts...)
}
