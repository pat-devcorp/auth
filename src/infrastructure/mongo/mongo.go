package mongo

import (
    "context"
    "fmt"
    "log"
    "os"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type MongoServerConfig struct {
    Hostname     string
    Port         int
    User         string
    Password     string
    DatabaseName string
    Collection   string
}

type MongoClient struct {
    client *mongo.Client
    dbName string
    collection *mongo.Collection
}

// NewClient creates a MongoClient instance
func NewClient(config MongoServerConfig) (*MongoClient, error) {
    uri := fmt.Sprintf("mongodb://%s:%s@%s:%d",
        config.User, config.Password, config.Hostname, config.Port)

    clientOptions := options.Client().ApplyURI(uri)
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        return nil, err
    }

    if err := client.Ping(context.Background(), nil); err != nil {
        return nil, err
    }

    return &MongoClient{
        client: client,
        dbName: config.DatabaseName,
        collection: client.Database(config.DatabaseName).Collection(config.Collection),
    }, nil
}

// SetDefaultFromEnv creates a MongoClient with default values from environment variables
func SetDefault(env *bootstrap.Env, tableName, pk string) (*MongoClient, error) {
    config := MongoServerConfig{
        Hostname:     env.MongoHost,
        Port:         env.MongoPort,
        User:         env.MongoUser,
        Password:     env.MongoPass,
        DatabaseName: env.MongoDB,
        Collection:   tableName,
    }
    return NewClient(config)
}

// Fetch fetches documents from the collection
func (mc *MongoClient) Fetch(ctx context.Context, filter interface{}) (*mongo.Cursor, error) {
    return mc.collection.Find(ctx, filter)
}

// Create inserts a new document into the collection
func (mc *MongoClient) Create(ctx context.Context, document interface{}) (string, error) {
    result, err := mc.collection.InsertOne(ctx, document)
    if err != nil {
        return "", err
    }
    return result.InsertedID.(string), nil
}

// GetById retrieves a document by its ID
func (mc *MongoClient) GetById(ctx context.Context, id string) (interface{}, error) {
    var result interface{}
    err := mc.collection.FindOne(ctx, bson.M{"_id": bson.M{"$eq": id}}).Decode(&result)
    if err == mongo.ErrNoDocuments {
        return nil, nil // Handle document not found gracefully
    }
    return result, err
}

// Update updates an existing document by its ID
func (mc *MongoClient) Update(ctx context.Context, id string, update interface{}) error {
    result, err := mc.collection.UpdateOne(ctx, bson.M{"_id": bson.M{"$eq": id}}, update)
    if err != nil {
        return err
    }
    if result.MatchedCount == 0 {
        return fmt.Errorf("document with id %s not found", id) // Handle document not found gracefully
    }
    return nil
}

// Delete deletes a document by its ID
func (mc *MongoClient) Delete(ctx context.Context, id string) error {
    result, err := mc.collection.DeleteOne(ctx, bson.M{"_id": bson.M{"$eq": id}})
    if err != nil {
        return err
    }
    if result.DeletedCount == 0 {
        return fmt.Errorf("document with id %s not found", id) // Handle document not found gracefully
    }
    return nil
}
