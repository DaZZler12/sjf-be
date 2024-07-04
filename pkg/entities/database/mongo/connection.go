package mongo

import (
	"context"
	"errors"
	"fmt"
	"sync"

	mongoORM "github.com/DaZZler12/sjf-be/pkg/entities/database/mongo/orm"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var (
	mongoDBInstance *mongo.Database
	once            sync.Once
	initErr         error
)

//here the mongoDBInstance is initialized only once,
//regardless of how many goroutines attempt to initialize it concurrently.

// The use of sync.Once provides a thread-safe way to enforce the Singleton pattern,
// ensuring that only one instance of the MongoDB client
// is created and shared across the application.

func GetMongoDBInstance(dbConfig *DatabaseConfig) (*mongo.Database, error) {
	once.Do(func() {
		if dbConfig == nil {
			initErr = errors.New("database config is nil, failed to connect to MongoDB")
			return
		}
		mongoDBInstance, initErr = ConnectToMongo(dbConfig)
	})
	return mongoDBInstance, initErr
}

// ConnectToMongo establishes a connection to MongoDB using the provided configuration.
func ConnectToMongo(dbConfig *DatabaseConfig) (*mongo.Database, error) {
	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%d", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port)
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}
	if err := client.Ping(context.TODO(), nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	zap.S().Info("Connected to MongoDB....")
	return client.Database(dbConfig.DBName), nil
}

// DisconnectFromMongo disconnects from MongoDB.
func DisconnectFromMongo(ctx context.Context) error {
	if mongoDBInstance == nil {
		return errors.New("mongoDBInstance is nil, failed to disconnect from MongoDB")
	}
	if err := mongoDBInstance.Client().Disconnect(ctx); err != nil {
		return fmt.Errorf("failed to disconnect from MongoDB: %w", err)
	}
	zap.S().Info("Disconnected from MongoDB....")
	return nil
}

func GetMongoCollection(collectionName string, opts ...*options.CollectionOptions) mongoORM.MongoCollectionDerived {
	if mongoDBInstance == nil {
		return nil
	}
	col := mongoDBInstance.Collection(collectionName, opts...)
	return &mongoORM.MongoORM{
		Col:     col,
		ColName: collectionName,
		Logger:  zap.L().Named(fmt.Sprintf("collection_%s", collectionName)).With(zap.String("collection", collectionName)),
	}
}
