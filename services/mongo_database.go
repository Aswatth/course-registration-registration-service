package services

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatabase struct {
	client       mongo.Client
	databse_name string
}

func (db *MongoDatabase) Connect(context context.Context, connection_string string) error {
	clientOptions := options.Client().ApplyURI(connection_string)

	client, err := mongo.Connect(context, clientOptions)
	db.client = *client
	return err
}

func (db *MongoDatabase) Disconnect(context context.Context) error {
	err := db.client.Disconnect(context)
	return err
}

func (db *MongoDatabase) Ping(context context.Context) error {
	err := db.client.Ping(context, nil)
	return err
}

func (db *MongoDatabase) SetDatabase(MongoDatabase_name string) {
	db.databse_name = MongoDatabase_name
}

func (db *MongoDatabase) GetCollection(collection_name string) (context.Context, mongo.Collection) {
	return context.Background(), *db.client.Database(db.databse_name).Collection(collection_name)
}

func (db *MongoDatabase) CreateCollection(collection_name string) {
	db.client.Database(db.databse_name).CreateCollection(context.Background(), collection_name)
}
