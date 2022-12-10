package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var TestCollection *mongo.Collection
var CTX = context.TODO()

func Init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(CTX, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(CTX, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Connected MongoDB")

	TestCollection = client.Database("20221210_go").Collection("test")
}
