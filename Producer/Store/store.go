package store

import (
	"time"
)

type HistoricalDataValues struct {
	Object string
	Price  int
}

type HistoricalData map[time.Time]HistoricalDataValues //This stores the historical transactions

// New - creates a new instance of the historical data
func New() HistoricalData {
	newInstance := make(HistoricalData)
	return newInstance
}

// Add - adds a new trade to the historical data
func (h HistoricalData) Add(newTradeTime time.Time, newTradeValues HistoricalDataValues) {
	h[newTradeTime] = newTradeValues
}
