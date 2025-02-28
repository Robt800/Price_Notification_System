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

// inMemoryStore - type to store the historical data privately.  This encapsulated map (within the struct) is used to
// facilitate easier unit testing.  Because inMemory implements the Storage interface, mocks can be injected more easily
// into the code
type inMemoryStore struct {
	data map[time.Time]HistoricalDataValues
}

// Add - adds a new trade to the historical data - i.e. the global memory store
func (h HistoricalData) Add(newTradeTime time.Time, newTradeValues HistoricalDataValues) {
	h[newTradeTime] = newTradeValues
}

// Add - adds a new trade to the historical data - i.e. the private memory store used to facilitate easier testing
func (s inMemoryStore) Add(newTradeTime time.Time, newTradeValues HistoricalDataValues) {
	s.data[newTradeTime] = newTradeValues
}
