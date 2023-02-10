package database

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoClient *mongo.Client
var err error

func Connect() {
	godotenv.Load()
	mongo_uri := os.Getenv("MONGODB_URI")

	MongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mongo_uri))

	if err != nil {
		panic(err)
	}

	// defer func() {
	// 	if err = MongoClient.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	if err := MongoClient.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("✅ Connected to MongoDB Successfully!")
}
