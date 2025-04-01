package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client *mongo.Client
}

func indexExists(collection *mongo.Collection, indexName string) (bool, error) {
	indexes, err := collection.Indexes().List(context.TODO())
	if err != nil {
		return false, err
	}
	defer indexes.Close(context.TODO())

	for indexes.Next(context.TODO()) {
		var index bson.M
		if err := indexes.Decode(&index); err != nil {
			return false, err
		}

		// Check if the index name matches the desired name
		if index["name"] == indexName {
			return true, nil
		}
	}

	// Check for errors that occurred during iteration
	if err := indexes.Err(); err != nil {
		return false, err
	}

	return false, nil
}

func NewMongoDB(uri string) *MongoDB {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	usersColl := client.Database("maincluster").Collection("users")

	if exists, _ := indexExists(usersColl, "location_2dsphere"); !exists {
		fmt.Println("Creating index on users.location")

		indexModel := mongo.IndexModel{
			Keys:    bson.D{{Key: "location", Value: "2dsphere"}},
			Options: &options.IndexOptions{},
		}

		name, err := usersColl.Indexes().CreateOne(context.TODO(), indexModel)

		if err != nil {
			panic("Failed to create index: " + err.Error())
		} else {
			fmt.Println("Name of Index Created: " + name)
		}
	}

	return &MongoDB{Client: client}
}

func (m *MongoDB) Close() {
	if err := m.Client.Disconnect(context.Background()); err != nil {
		log.Fatalf("Error disconnecting from MongoDB: %v", err)
	}
}
