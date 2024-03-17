package mongodb_dal

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoDBDAL struct {
	client     *mongo.Client
	database   *mongo.Database
	Collection *mongo.Collection
}

func NewMongoDBDAL(connectionString, dbName, collectionName string) (*MongoDBDAL, error) {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	db := client.Database(dbName)
	coll := db.Collection(collectionName)

	return &MongoDBDAL{
		client:     client,
		database:   db,
		Collection: coll,
	}, nil
}

func (d *MongoDBDAL) InsertOne(data interface{}) (*mongo.InsertOneResult, error) {
	result, err := d.Collection.InsertOne(context.Background(), data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (d *MongoDBDAL) FindOne(filter interface{}, result interface{}) error {
	err := d.Collection.FindOne(context.Background(), filter).Decode(result)
	if err != nil {
		return err
	}
	return nil
}

func (d *MongoDBDAL) UpdateOne(filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	result, err := d.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (d *MongoDBDAL) DeleteOne(filter interface{}) (*mongo.DeleteResult, error) {
	result, err := d.Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DropDatabase drops the entire database
func (d *MongoDBDAL) DropDatabase() error {
	err := d.client.Database(d.database.Name()).Drop(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (d *MongoDBDAL) Close() {
	d.client.Disconnect(context.Background())
}
