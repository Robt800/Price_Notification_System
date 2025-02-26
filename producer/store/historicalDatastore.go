package store

import (
	"time"
)

type Storage interface {
	Add(time.Time, HistoricalDataValues)
}

type HistoricalDataValues struct {
	Object string
	Price  int
}

// HistoricalData - This stores the historical transactions
type HistoricalData map[time.Time]HistoricalDataValues

// ItemTradeHistory - a shared data store for the historical transactions
var ItemTradeHistory = make(HistoricalData)

// Add - adds a new trade to the historical data
func (h HistoricalData) Add(newTradeTime time.Time, newTradeValues HistoricalDataValues) {
	h[newTradeTime] = newTradeValues
}
