package store

import (
	"Price_Notification_System/models"
)

// inMemoryAlertStore - concrete implementation of the AlertDefStore interface
type inMemoryAlertStore struct {
	data map[string]models.AlertValues
}

// compile time check to ensure that inMemoryTradeStore implements the TradeStore interface
// This won't be used within the code - but if not implemented correctly - the compiler will throw an error
var _ AlertDefStore = &inMemoryAlertStore{}

func NewInMemoryAlertStore() AlertDefStore {
	return &inMemoryAlertStore{
		data: make(map[string]models.AlertValues),
	}
}

// AddAlert - adds a new alert to the alerts active - i.e. the private memory store used to facilitate easier testing
func (i *inMemoryAlertStore) AddAlert(itemToAlert string, newAlertDef models.AlertValues) {
	i.data[itemToAlert] = newAlertDef
}
