package service

import (
	"Price_Notification_System/models"
	"Price_Notification_System/store"
	"context"
)

func GetTradesByItem(
	ctx context.Context,
	tradeStore store.TradeStore,
	item string,
) (data []models.HistoricalTradeDataReturned, err error) {
	data, err = tradeStore.GetTradeByItem(item)
	return data, err
}

func ProcessAlerts(
	ctx context.Context,
	alertStore store.AlertDefStore,
	tradeStore store.TradeStore,
	item string,
) (alertActive bool, data []models.HistoricalTradeAlertReturned, err error) {

	// Get specific alerts for the item
	alerts, errFromGetAlertsByItem := alertStore.GetAlertsByItem(item)
	if errFromGetAlertsByItem != nil {
		return false, nil, errFromGetAlertsByItem
	}

	//Get specific trades for the item
	returnedData, errFromGetTradeByItem := tradeStore.GetTradeByItem(item)
	if errFromGetTradeByItem != nil {
		return false, nil, errFromGetTradeByItem
	}

	// iterate over the returned data and check if the price is alertable
	for _, itemTrades := range returnedData {
		for _, alertByItem := range alerts {
			if alertByItem.AlertType == models.PriceAlertLowPrice {
				if itemTrades.Price < alertByItem.PriceTrigger {
					data = append(data, models.HistoricalTradeAlertReturned{HistoricalTradeDataReturned: itemTrades,
						AlertValues: models.AlertValues{
							AlertType: alertByItem.AlertType, PriceTrigger: alertByItem.PriceTrigger}})
				}
			}
			if alertByItem.AlertType == models.PriceAlertHighPrice {
				if itemTrades.Price > alertByItem.PriceTrigger {
					data = append(data, models.HistoricalTradeAlertReturned{HistoricalTradeDataReturned: itemTrades,
						AlertValues: models.AlertValues{
							AlertType: alertByItem.AlertType, PriceTrigger: alertByItem.PriceTrigger}})
				}
			}
		}
	}

	// Determine the return values
	if len(data) > 0 {
		return true, data, nil
	} else {
		return false, nil, models.ErrNoDataMatchingAlertsFound
	}

}

func CreateNewAlert(
	ctx context.Context,
	alertStore store.AlertDefStore,
	item string,
	newAlertDef models.AlertValues,
) (err error) {

	// Add the alert to the alert store
	err = alertStore.AddAlert(item, newAlertDef)
	if err != nil {
		return err
	}

	return nil
}
