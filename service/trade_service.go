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
	alertToProcess models.AlertDef,
) (data []models.HistoricalTradeAlertReturned, err error) {

	// iterate over the alerts and get the trade data that matches the alert item
	for item, alertDef := range alertStore.data {
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
