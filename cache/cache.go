// In your data/cache.go file

package data

import (
	"context"
	"sync"
	"time"

	"github.com/atul-007/stockPriceView/data"
	"github.com/atul-007/stockPriceView/models"
)

type CacheEntry struct {
	Expiration time.Time
	Data       interface{}
}

// Cache represents an in-memory cache for stock price history
type Cache struct {
	mu    sync.RWMutex
	cache map[string]CacheEntry
}

// NewCache creates a new instance of the cache
func NewCache() *Cache {
	return &Cache{
		cache: make(map[string]CacheEntry),
	}
}

// Get retrieves a value from the cache based on a key
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, ok := c.cache[key]
	if !ok || entry.Expiration.Before(time.Now()) {
		return nil, false
	}

	return entry.Data, true
}

// Set adds or updates a value in the cache based on a key
func (c *Cache) Set(key string, data interface{}, expiration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = CacheEntry{
		Expiration: time.Now().Add(expiration),
		Data:       data,
	}
}

func GetStockPriceHistoryWithCache(ctx context.Context, cache *Cache, code string) ([]models.StockPriceHistory, error) {
	// Try to get data from the cache
	if cachedData, ok := cache.Get(code); ok {
		if priceHistory, ok := cachedData.([]models.StockPriceHistory); ok {
			return priceHistory, nil
		}
	}

	// If not in the cache, fetch data from the database
	priceHistory, err := data.GetStockPriceHistory(ctx, code)
	if err != nil {
		return nil, err
	}

	// Convert data.StockPriceHistory to models.StockPriceHistory
	var convertedPriceHistory []models.StockPriceHistory
	for _, entry := range priceHistory {
		convertedEntry := models.StockPriceHistory{
			Date:  entry.Date,
			Price: entry.Price,
			// Add other fields as needed
		}
		convertedPriceHistory = append(convertedPriceHistory, convertedEntry)
	}

	// Set the fetched data in the cache with a specified expiration (adjust as needed)
	cache.Set(code, convertedPriceHistory, 5*time.Minute)

	return convertedPriceHistory, nil
}
