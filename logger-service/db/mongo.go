package db

import (
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var Client *mongo.Client

func InitDb() {
	connStr := os.Getenv("mongoDB")
	mongoClient, err := mongo.Connect(options.Client().ApplyURI(connStr))
	if err != nil {
		panic("Unable to connect to the database.")
	}
	Client = mongoClient
}
