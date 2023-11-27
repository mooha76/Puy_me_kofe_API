package utils

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CollectionExists(ctx context.Context, coll *mongo.Collection) bool {
	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		return false
	}
	defer cursor.Close(ctx)
	return cursor.Next(ctx)
}

func DatabaseExists(ctx context.Context, DB *mongo.Database) bool {
	cursor, err := DB.ListCollections(ctx, bson.M{})
	if err != nil {
		return false
	}
	defer cursor.Close(ctx)
	return cursor.Next(ctx)

}
