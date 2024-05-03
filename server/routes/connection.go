package routes

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client *mongo.Client = DBInstance()
)

func DBInstance()  *mongo.Client{
	MongoDb := "mongodb://root:developer@localhost:27017/caloriesdb?authSource=admin"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoDb))

	if err != nil {
		log.Fatal(err)
	}
	defer cancel()
	// err = client.Connect(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println("Connected to MongoDB")
	return client
}

func openCollection(client *mongo.Client, collectionName string)  *mongo.Collection {
	var collection *mongo.Collection = client.Database("caloriesdb").Collection(collectionName)
	return collection
}