package mongo

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/mooha76/Puy_me_kofe_API/config"
	utils "github.com/mooha76/Puy_me_kofe_API/utils/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Mongoconnection() (config.Config, error) {
	// Set your MongoDB connection string
	mongoConnectionString := "mongodb://localhost:27017"
	// Set your database and collection names
	dbName := "Gomo"
	collectionName := "config"

	// Set a context with timeout
	ctx := context.Background()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoConnectionString))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx) // Close the MongoDB connection when done

	// Select database and collection
	db := client.Database(dbName)
	collection := db.Collection(collectionName)

	// Check if the database exists
	if !utils.DatabaseExists(ctx, db) {
		return config.Config{}, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Check if the collection exists
	if !utils.CollectionExists(ctx, collection) {
		return config.Config{}, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Query for all documents in the collection
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return config.Config{}, err
	}
	defer cursor.Close(ctx)

	// Assuming you want to return the first document found
	if cursor.Next(ctx) {
		var result config.Config
		err := cursor.Decode(&result)
		if err != nil {
			return config.Config{}, err
		}
		return result, nil
	}

	return config.Config{}, errors.New("No documents found in the collection")
}
