package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ArchiMrin/Project_GOLang/eCOMM/constants"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDataBase() (*mongo.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mongoConnection := options.Client().ApplyURI(constants.ConnectionString)
	mongoClient, err := mongo.Connect(ctx, mongoConnection)
	if err != nil {
		log.Fatal(err.Error())
	}
	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	fmt.Println("Connected to DB")
	return mongoClient, nil
}

func GetCollection(client *mongo.Client, dbName string, collectionName string) *mongo.Collection {
	collection := client.Database(dbName).Collection(collectionName)
	return collection
}
