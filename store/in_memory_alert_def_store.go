package store

import (
	"Price_Notification_System/models"
)

// inMemoryAlertStore - concrete implementation of the AlertDefStore interface
type InMemoryAlertStore struct {
	data map[string]models.AlertValues
}

// compile time check to ensure that inMemoryTradeStore implements the TradeStore interface
// This won't be used within the code - but if not implemented correctly - the compiler will throw an error
var _ AlertDefStore = &InMemoryAlertStore{}

func NewInMemoryAlertStore() AlertDefStore {
	return &InMemoryAlertStore{
		data: make(map[string]models.AlertValues),
	}
}

// AddAlert - adds a new alert to the alerts active - i.e. the private memory store used to facilitate easier testing
func (i *InMemoryAlertStore) AddAlert(itemToAlert string, newAlertDef models.AlertValues) {
	i.data[itemToAlert] = newAlertDef
}

// GetAlertsByItem - retrieves the alerts for a specific item
func (i *InMemoryAlertStore) GetAlertsByItem(item string) (data []models.AlertsByItemReturned, err error) {
	var matchingData models.AlertsByItemReturned
	for itemFromMemory, alertVals := range i.data {

		if itemFromMemory == item {
			matchingData = models.AlertsByItemReturned{
				Item: itemFromMemory,
				AlertValues: models.AlertValues{
					AlertType:    alertVals.AlertType,
					PriceTrigger: alertVals.PriceTrigger},
			}

			data = append(data, matchingData)
		}
	}
	if len(data) > 0 {
		return data, nil
	} else {
		return nil, models.ErrNoAlertsForItemFound
	}
}
