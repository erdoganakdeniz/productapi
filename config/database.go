package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}
type MongoOptions struct {
	New options.FindOneAndUpdateOptions
}

var (
	Mongo    MongoInstance
	MongoOps MongoOptions
)

func SetupDB() {
	dbName := os.Getenv("DB_NAME")
	mongoURI := os.Getenv("DB_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	db := client.Database(dbName)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	Mongo = MongoInstance{
		Client: client,
		DB:     db,
	}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	MongoOps = MongoOptions{
		New: opt,
	}
}
