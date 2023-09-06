package database

import (
	"context"
	"go-api-insta/helpers/variable"
	"go-api-insta/libs/logger"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Init() {

	var ctx = context.TODO()

	clientOptions := options.Client().ApplyURI(variable.GetEnvVariable("MONGODB_URI"))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	logger.Production.Info("Connected to Mongo database")

}

func Dbconnect() *mongo.Client {

	var ctx = context.Background()

	clientOptions := options.Client().ApplyURI(variable.GetEnvVariable("MONGODB_URI"))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return client

}
