package api

import (
	"encoding/json"
	"net/http"

	"github.com/atul-007/stockPriceView/data"
	"github.com/atul-007/stockPriceView/models"
)

func handleError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func Health(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Server Is up and running")

}

// GetTop10StocksHandler handles the request for getting the top 10 stocks
func GetTop10StocksHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Fetch top 10 stocks
	stocks, err := data.GetTop10Stocks(ctx, data.GetMongoDBCollection())
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Failed to fetch top 10 stocks")
		return
	}

	// Respond with the top 10 stocks in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stocks)
}

// GetStockByNameHandler handles the request for getting a stock by name
func GetStockByNameHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Extract the stock name from the query parameters
	name := r.URL.Query().Get("name")
	if name == "" {
		handleError(w, http.StatusBadRequest, "Stock name parameter is missing")
		return
	}
	// Get the stock by name
	stock, err := data.GetStockByName(ctx, data.GetMongoDBCollection(), name)
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Failed to fetch stock by name")
		return
	}

	// Respond with the stock details in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stock)
}

// GetStockPriceHistoryHandler handles the request for getting stock price history
func GetStockPriceHistoryHandler(w http.ResponseWriter, r *http.Request) {
	// Get the stock code from the query parameters
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Missing 'code' parameter", http.StatusBadRequest)
		return
	}

	// Get stock price history from the database
	priceHistory, err := data.GetStockPriceHistory(r.Context(), code)
	if err != nil {
		http.Error(w, "Error retrieving stock price history", http.StatusInternalServerError)
		return
	}

	// Return the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(priceHistory)
}

// AddToFavoritesHandler handles the request for adding a stock to favorites
func AddToFavoritesHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Parse the request body to get the stock details
	var stock models.Stock
	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		handleError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Add the stock to favorites
	err = data.AddToFavorites(ctx, data.GetMongoDBCollection(), stock)
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Failed to add stock to favorites")
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusNoContent)
}

// GetFavoriteStocksHandler handles the request for getting favorite stocks
func GetFavoriteStocksHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Get favorite stocks
	favoriteStocks, err := data.GetFavoriteStocks(ctx, data.GetMongoDBCollection())
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Failed to fetch favorite stocks")
		return
	}

	// Respond with favorite stocks in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(favoriteStocks)
}

// RemoveFromFavoritesHandler handles the request for removing a stock from favorites
func RemoveFromFavoritesHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Extract the stock code from the query parameters
	code := r.URL.Query().Get("code")
	if code == "" {
		handleError(w, http.StatusBadRequest, "Stock code parameter is missing")
		return
	}

	// Remove the stock from favorites
	err := data.RemoveFromFavorites(ctx, data.GetMongoDBCollection(), code)
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Failed to remove stock from favorites")
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusNoContent)
}
