package store

import (
	"Price_Notification_System/models"
)

// InMemoryAlertStore - concrete implementation of the AlertDefStore interface
type InMemoryAlertStore struct {
	data []models.AlertDef
}

// compile time check to ensure that inMemoryTradeStore implements the TradeStore interface
// This won't be used within the code - but if not implemented correctly - the compiler will throw an error
var _ AlertDefStore = &InMemoryAlertStore{}

func NewInMemoryAlertStore() AlertDefStore {
	return &InMemoryAlertStore{
		data: make([]models.AlertDef, 0, 20),
	}
}

func NewInMemoryAlertStoreWithData(data *[]models.AlertDef) AlertDefStore {
	return &InMemoryAlertStore{
		data: *data,
	}
}

// AddAlert - adds a new alert to the alerts active - i.e. the private memory store used to facilitate easier testing
func (i *InMemoryAlertStore) AddAlert(itemToAlert string, newAlertDef models.AlertValues) {
	//i.data[it = newAlertDef
	i.data = append(i.data, models.AlertDef{Item: itemToAlert, AlertValues: newAlertDef})
}

// GetAlertsByItem - retrieves the alerts for a specific item
func (i *InMemoryAlertStore) GetAlertsByItem(item string) (data []models.AlertsByItemReturned, err error) {
	var matchingData models.AlertsByItemReturned
	for _, v := range i.data {

		if v.Item == item {
			matchingData = models.AlertsByItemReturned{
				Item: v.Item,
				AlertValues: models.AlertValues{
					AlertType:    v.AlertType,
					PriceTrigger: v.PriceTrigger},
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
