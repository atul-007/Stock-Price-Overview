package data

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/atul-007/stockPriceView/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SaveToMongoDB reads the CSV file and saves the data to MongoDB using the provided client
func SaveToMongoDB(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	header, err := reader.Read()
	if err != nil {
		return err
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoDBConnectionString))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return err
	}
	defer client.Disconnect(ctx)

	collection := client.Database(databaseName).Collection(collectionName)

	var stocks []interface{}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		stock := models.Stock{}
		for i, value := range record {
			// Assuming the order of columns in the CSV file is fixed
			switch header[i] {
			case "SC_CODE":
				stock.Code = value
			case "SC_NAME":
				stock.Name = value
			case "SC_GROUP":
				stock.Group = value
			case "SC_TYPE":
				stock.Type = value
			case "OPEN":
				stock.Open, _ = strconv.ParseFloat(value, 64)
			case "HIGH":
				stock.High, _ = strconv.ParseFloat(value, 64)
			case "LOW":
				stock.Low, _ = strconv.ParseFloat(value, 64)
			case "CLOSE":
				stock.Close, _ = strconv.ParseFloat(value, 64)
			case "LAST":
				stock.Last, _ = strconv.ParseFloat(value, 64)
			case "PREVCLOSE":
				stock.PrevClose, _ = strconv.ParseFloat(value, 64)
			case "NO_TRADES":
				stock.NoTrades, _ = strconv.Atoi(value)
			case "NO_OF_SHRS":
				stock.NoOfShares, _ = strconv.Atoi(value)
			case "NET_TURNOV":
				stock.NetTurnover, _ = strconv.ParseFloat(strings.Trim(value, ","), 64)
			}
		}

		// Check if the stock already exists in the collection
		existingStock, err := GetStockByCode(ctx, collection, stock.Code)
		if err != nil && err != mongo.ErrNoDocuments {
			return err
		}

		if existingStock.Code == "" {
			// If the stock does not exist, insert it
			stocks = append(stocks, stock)
		} else {
			// If the stock already exists, update it (you can implement your update logic here)
			// For simplicity, we'll print a message here
			fmt.Printf("Stock with code '%s' already exists. Updating...\n", stock.Code)
		}
	}

	// Insert or update data into MongoDB
	for _, stock := range stocks {
		filter := bson.D{{Key: "code", Value: stock.(models.Stock).Code}}
		update := bson.D{{Key: "$set", Value: stock}}

		_, err := collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
		if err != nil {
			return err
		}
	}

	// Example: Query the top 10 stocks
	top10Stocks, err := GetTop10Stocks(ctx, collection)
	if err != nil {
		return err
	}
	fmt.Println("Top 10 Stocks:")
	for _, stock := range top10Stocks {
		fmt.Printf("%s - %s\n", stock.Code, stock.Name)
	}

	// Example: Query stocks by name
	nameQuery := "ABB LTD."
	abbStock, err := GetStockByName(ctx, collection, nameQuery)
	if err != nil {
		return err
	}
	fmt.Printf("Stock with Name '%s': %+v\n", nameQuery, abbStock)

	// Example: Add a stock to favorites
	favoriteStock := models.Stock{
		Code: "XYZ123",
		Name: "Favorite Stock",
		// Add other fields as needed
	}
	err = AddToFavorites(ctx, collection, favoriteStock)
	if err != nil {
		return err
	}
	fmt.Printf("Stock '%s' added to favorites\n", favoriteStock.Code)

	// Example: View favorite stocks
	favoriteStocks, err := GetFavoriteStocks(ctx, collection)
	if err != nil {
		return err
	}
	fmt.Println("Favorite Stocks:")
	for _, favStock := range favoriteStocks {
		fmt.Printf("%s - %s\n", favStock.Code, favStock.Name)
	}

	// Example: Remove a stock from favorites
	err = RemoveFromFavorites(ctx, collection, favoriteStock.Code)
	if err != nil {
		return err
	}
	fmt.Printf("Stock '%s' removed from favorites\n", favoriteStock.Code)

	return nil
}

// GetTop10Stocks retrieves the top 10 stocks from the MongoDB collection
func GetTop10Stocks(ctx context.Context, collection *mongo.Collection) ([]models.Stock, error) {
	var top10Stocks []models.Stock

	limit := int64(10) // Convert 10 to int64

	cursor, err := collection.Find(ctx, bson.D{}, &options.FindOptions{
		Sort:  bson.D{{Key: "close", Value: -1}},
		Limit: &limit, // Pass the address of the int64 variable
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &top10Stocks)
	if err != nil {
		return nil, err
	}

	return top10Stocks, nil
}

// GetStockByName retrieves a stock by name from the MongoDB collection
func GetStockByName(ctx context.Context, collection *mongo.Collection, name string) (models.Stock, error) {
	var stock models.Stock

	filter := bson.D{{Key: "name", Value: name}}

	err := collection.FindOne(ctx, filter).Decode(&stock)
	if err != nil {
		return models.Stock{}, err
	}

	return stock, nil
}

// AddToFavorites adds a stock to the favorites in the MongoDB collection
func AddToFavorites(ctx context.Context, collection *mongo.Collection, stock models.Stock) error {
	_, err := collection.InsertOne(ctx, stock)
	return err
}

// GetFavoriteStocks retrieves favorite stocks from the MongoDB collection
func GetFavoriteStocks(ctx context.Context, collection *mongo.Collection) ([]models.Stock, error) {
	var favoriteStocks []models.Stock

	cursor, err := collection.Find(ctx, bson.D{{Key: "favorite", Value: true}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &favoriteStocks)
	if err != nil {
		return nil, err
	}

	return favoriteStocks, nil
}

// RemoveFromFavorites removes a stock from favorites in the MongoDB collection
func RemoveFromFavorites(ctx context.Context, collection *mongo.Collection, code string) error {
	filter := bson.D{{Key: "code", Value: code}}

	_, err := collection.DeleteOne(ctx, filter)
	return err
}
func GetStockPriceHistory(ctx context.Context, code string) ([]models.StockPriceHistory, error) {
	var priceHistory []models.StockPriceHistory

	collection := GetMongoDBCollection() // Implement this function based on your database connection

	filter := bson.D{{Key: "code", Value: code}}
	projection := bson.D{{Key: "_id", Value: 0}, {Key: "date", Value: 1}, {Key: "close", Value: 1}}

	cursor, err := collection.Find(ctx, filter, &options.FindOptions{Projection: projection})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &priceHistory)
	if err != nil {
		return nil, err
	}

	return priceHistory, nil
}
