package store

import (
	"Price_Notification_System/models"
)

// inMemoryAlertStore - concrete implementation of the AlertDefStore interface
type inMemoryAlertStore struct {
	data       map[string]models.AlertValues
	tradeStore TradeStore
}

// compile time check to ensure that inMemoryTradeStore implements the TradeStore interface
// This won't be used within the code - but if not implemented correctly - the compiler will throw an error
var _ AlertDefStore = &inMemoryAlertStore{}

func NewInMemoryAlertStore(ts TradeStore) AlertDefStore {
	return &inMemoryAlertStore{
		data:       make(map[string]models.AlertValues),
		tradeStore: ts,
	}
}

// AddAlert - adds a new alert to the alerts active - i.e. the private memory store used to facilitate easier testing
func (i *inMemoryAlertStore) AddAlert(itemToAlert string, newAlertDef models.AlertValues) {
	i.data[itemToAlert] = newAlertDef
}

func (i *inMemoryAlertStore) ProcessAlerts(alertToProcess models.AlertDef) (data []models.HistoricalTradeAlertReturned, err error) {

	// iterate over the alerts and get the trade data that matches the alert item
	for item, alertDef := range i.data {
		returnedData, errFromGetTradeByItem := i.tradeStore.GetTradeByItem(item)
		if errFromGetTradeByItem != nil {
			return nil, errFromGetTradeByItem
		}

		// iterate over the returned data and check if the price is alertable
		if alertDef.AlertType == models.PriceAlertLowPrice {
			for _, v := range returnedData {
				if v.Price < alertDef.PriceTrigger {
					data = append(data, models.HistoricalTradeAlertReturned{HistoricalTradeDataReturned: v,
						AlertValues: models.AlertValues{
							AlertType: alertDef.AlertType, PriceTrigger: alertDef.PriceTrigger,
						},
					})
				}
			}
		} else if alertDef.AlertType == models.PriceAlertHighPrice {
			for _, v := range returnedData {
				if v.Price > alertDef.PriceTrigger {
					data = append(data, models.HistoricalTradeAlertReturned{HistoricalTradeDataReturned: v,
						AlertValues: models.AlertValues{
							AlertType: alertDef.AlertType, PriceTrigger: alertDef.PriceTrigger,
						},
					})
				}
			}
		}

	}
	if len(data) > 0 {
		return data, nil
	} else {
		return nil, models.ErrNoDataMatchingAlertsFound
	}
}
