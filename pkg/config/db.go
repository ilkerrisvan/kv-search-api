package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func DBConnect() *mongo.Client {
	URL := os.Getenv("DB_URL")
	if URL == "" {
		log.Fatal("err")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(URL))
	if err != nil {
		panic(err)
	}

	return client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("getircase-study").Collection(collectionName)
}
