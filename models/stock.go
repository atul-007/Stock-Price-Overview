package models

import (
	"sync"
	"time"
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
type StockPriceHistory struct {
	Date  time.Time `bson:"date"`
	Price float64   `bson:"price"`
}
type Stock struct {
	Code        string  `bson:"code"`
	Name        string  `bson:"name"`
	Group       string  `bson:"group"`
	Type        string  `bson:"type"`
	Open        float64 `bson:"open"`
	High        float64 `bson:"high"`
	Low         float64 `bson:"low"`
	Close       float64 `bson:"close"`
	Last        float64 `bson:"last"`
	PrevClose   float64 `bson:"prev_close"`
	NoTrades    int     `bson:"no_trades"`
	NoOfShares  int     `bson:"no_of_shares"`
	NetTurnover float64 `bson:"net_turnover"`
}
