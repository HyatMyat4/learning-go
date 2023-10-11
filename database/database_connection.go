package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBconnect() *mongo.Client {
	MongoURI := "mongodb://mongo:JHAucBRYLvUFsMWZoZAE@containers-us-west-201.railway.app:5666"
	fmt.Print(MongoURI)

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoURI))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("* MongoDB Connection Successfully *")

	return client
}

var Client *mongo.Client = DBconnect()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	var collection *mongo.Collection = client.Database("test").Collection(collectionName)

	return collection
}
