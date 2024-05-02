package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	MongoDB := os.Getenv("MONGO_URI")
	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(MongoDB).SetServerAPIOptions(serverApi)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatalf("Error creating mongo client: %v", err)
	}
	if err := client.Database("restaurant").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		log.Fatalf("Error pinging mongo client: %v", err)
	}
	fmt.Println("Connected to MongoDB!✅")
	return client
}

var Client *mongo.Client = DBInstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("restaurant").Collection(collectionName)
	return collection
}
