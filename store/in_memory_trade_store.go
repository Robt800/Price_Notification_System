package store

import (
	"Price_Notification_System/models"
	"time"
)

// inMemoryTradeStore - concrete implementation of the TradeStore interface
type InMemoryTradeStore struct {
	TradeData map[time.Time]models.HistoricalDataValues
}

// compile time check to ensure that inMemoryTradeStore implements the TradeStore interface
// This won't be used within the code - but if not implemented correctly - the compiler will throw an error
var _ TradeStore = &InMemoryTradeStore{}

// NewInMemoryTradeStore - constructor function to create a new instance of the inMemoryTradeStore
func NewInMemoryTradeStore() TradeStore {
	return &InMemoryTradeStore{
		TradeData: make(map[time.Time]models.HistoricalDataValues),
	}
}

// NewInMemoryTradeStoreWithData - constructor function to create a new instance of the inMemoryTradeStore with data
func NewInMemoryTradeStoreWithData(data *map[time.Time]models.HistoricalDataValues) TradeStore {
	return &InMemoryTradeStore{
		TradeData: *data,
	}
}

func (i *InMemoryTradeStore) AddTrade(tradeTime time.Time, tradeValues models.HistoricalDataValues) {
	i.TradeData[tradeTime] = tradeValues
}

func (i *InMemoryTradeStore) GetTradeByItem(item string) (data []models.HistoricalTradeDataReturned, err error) {
	var matchingData models.HistoricalTradeDataReturned
	for tradeTime, tradeObjPrice := range i.TradeData {

		if tradeObjPrice.Object == item {
			matchingData = models.HistoricalTradeDataReturned{
				Date: tradeTime,
				HistoricalDataValues: models.HistoricalDataValues{
					Object: tradeObjPrice.Object,
					Price:  tradeObjPrice.Price},
			}

			data = append(data, matchingData)
		}
	}
	if len(data) > 0 {
		return data, nil
	} else {
		return nil, models.ErrTradeNotFound
	}
}
