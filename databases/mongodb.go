package databases

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)


func DB() (*mongo.Database,*mongo.Collection) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://Admin:pJS8u7TdIiaJwA7n@testcluster.kg1gr.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	database:=client.Database("ExampleDatabase")
	collection:=database.Collection("exampleCollection")
	return database,collection
}