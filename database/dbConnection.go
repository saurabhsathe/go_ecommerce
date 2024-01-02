package routes

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	mongoDB, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/restaurantdb"))

	if err != nil {
		log.Fatal(err)
	}
	defer cancel()
	return mongoDB
}

var Client *mongo.Client = DBInstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	mongoCollection := client.Database("restaurantdb").Collection(collectionName)
	return mongoCollection
}
