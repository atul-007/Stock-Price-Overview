package data

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/atul-007/stockPriceView/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoClient   *mongo.Client
	mongoClientMu sync.Mutex
)

// MongoDB connection string
const mongoDBConnectionString = "mongodb+srv://atulranjan789:atulranjan@cluster0.ebjtzlx.mongodb.net/?retryWrites=true&w=majority"

// MongoDB database name
const databaseName = "stock_prices"

// MongoDB collection name
const collectionName = "equity_bhavcopy"

// GetMongoDBCollection returns the MongoDB collection instance
func GetMongoDBCollection() *mongo.Collection {
	client := getMongoClient()
	return client.Database(databaseName).Collection(collectionName)
}

// InitializeMongoDBConnection initializes the MongoDB connection
func InitializeMongoDBConnection() (*mongo.Client, error) {
	client := getMongoClient()
	return client, nil
}

// getMongoClient returns a singleton MongoDB client instance
func getMongoClient() *mongo.Client {
	if mongoClient != nil {
		return mongoClient
	}

	mongoClientMu.Lock()
	defer mongoClientMu.Unlock()

	// Check again inside the lock to avoid race condition
	if mongoClient != nil {
		return mongoClient
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoDBConnectionString))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the MongoDB server to verify the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	mongoClient = client
	return mongoClient
}

func GetStockByCode(ctx context.Context, collection *mongo.Collection, code string) (models.Stock, error) {
	var stock models.Stock

	filter := bson.D{{Key: "code", Value: code}}

	err := collection.FindOne(ctx, filter).Decode(&stock)
	if err != nil {
		return models.Stock{}, err
	}

	return stock, nil
}
