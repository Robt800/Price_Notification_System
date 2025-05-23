package service

import (
	"Price_Notification_System/models"
	"Price_Notification_System/service"
	"Price_Notification_System/store"
	"context"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestTradeService_ProcessAlertsByItem(t *testing.T) {

	type testDefProcessAlertsByItem struct {
		instanceName        string
		alertStore          store.AlertDefStore
		tradeStore          store.TradeStore
		item                string
		expectedAlertActive bool
		expectedData        []models.HistoricalTradeAlertReturned
		expectedError       error
		wantErr             bool
	}

	tests := []testDefProcessAlertsByItem{
		{instanceName: "test1 - Process alerts for hulk item",
			alertStore: store.NewInMemoryAlertStoreWithData(&[]models.AlertDef{
				{Item: "hulk", AlertValues: models.AlertValues{AlertType: models.PriceAlertLowPrice, PriceTrigger: 975}},
				{Item: "wolverine", AlertValues: models.AlertValues{AlertType: models.PriceAlertLowPrice, PriceTrigger: 600}},
				{Item: "hulk", AlertValues: models.AlertValues{AlertType: models.PriceAlertLowPrice, PriceTrigger: 600}},
				{Item: "superman", AlertValues: models.AlertValues{AlertType: models.PriceAlertLowPrice, PriceTrigger: 1200}},
				{Item: "hulk", AlertValues: models.AlertValues{AlertType: models.PriceAlertHighPrice, PriceTrigger: 1075}},
			}),
			tradeStore: store.NewInMemoryTradeStoreWithData(&map[time.Time]models.HistoricalDataValues{
				time.Date(2022, 1, 1, 10, 20, 20, 0, time.UTC):  {Object: "hulk", Price: 1000},
				time.Date(2023, 4, 2, 6, 45, 14, 0, time.UTC):   {Object: "wolverine", Price: 1100},
				time.Date(2023, 5, 3, 8, 12, 18, 0, time.UTC):   {Object: "hulk", Price: 1200},
				time.Date(2023, 10, 4, 9, 19, 20, 0, time.UTC):  {Object: "superman", Price: 1300},
				time.Date(2024, 12, 5, 13, 46, 23, 0, time.UTC): {Object: "hulk", Price: 590},
				time.Date(2024, 2, 6, 18, 52, 56, 0, time.UTC):  {Object: "hulk", Price: 900},
			}),
			item: "hulk", expectedAlertActive: true,
			expectedData: []models.HistoricalTradeAlertReturned{
				{HistoricalTradeDataReturned: models.HistoricalTradeDataReturned{Date: time.Date(2023, 5, 3, 8, 12, 18, 0, time.UTC),
					HistoricalDataValues: models.HistoricalDataValues{Object: "hulk", Price: 1200}},
					AlertValues: models.AlertValues{AlertType: models.PriceAlertHighPrice, PriceTrigger: 1075}},
				{HistoricalTradeDataReturned: models.HistoricalTradeDataReturned{Date: time.Date(2024, 12, 5, 13, 46, 23, 0, time.UTC),
					HistoricalDataValues: models.HistoricalDataValues{Object: "hulk", Price: 590}},
					AlertValues: models.AlertValues{AlertType: models.PriceAlertLowPrice, PriceTrigger: 975}},
				{HistoricalTradeDataReturned: models.HistoricalTradeDataReturned{Date: time.Date(2024, 12, 5, 13, 46, 23, 0, time.UTC),
					HistoricalDataValues: models.HistoricalDataValues{Object: "hulk", Price: 590}},
					AlertValues: models.AlertValues{AlertType: models.PriceAlertLowPrice, PriceTrigger: 600}},
				{HistoricalTradeDataReturned: models.HistoricalTradeDataReturned{Date: time.Date(2024, 2, 6, 18, 52, 56, 0, time.UTC),
					HistoricalDataValues: models.HistoricalDataValues{Object: "hulk", Price: 900}},
					AlertValues: models.AlertValues{AlertType: models.PriceAlertLowPrice, PriceTrigger: 975}},
			},
			expectedError: nil, wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.instanceName, func(t *testing.T) {

			//Call the ProcessAlerts function with the test data
			ctx := context.Background()
			alertsActive, data, err := service.ProcessAlerts(ctx, tt.alertStore, tt.tradeStore, tt.item)

			//Sort the data returned by date - this is to ensure that the test is not dependent on the order of the data returned
			//Also sort the test data to ensure it is in the same order
			sortSliceByDate(&data)
			sortSliceByDate(&tt.expectedData)

			//Check if the alerts were processed correctly
			//Check if any errors were as expected
			if err != nil && !tt.wantErr {
				t.Errorf("ProcessAlerts() received error = %v, wanted %v", err, tt.expectedError)
			}
			if err == nil && tt.wantErr {
				t.Errorf("ProcessAlerts() received no error, wanted %v", tt.expectedError)
			}

			//Check if the alertsActive value is as expected
			if alertsActive != tt.expectedAlertActive {
				t.Errorf("ProcessAlerts() alertsActive = %v, wanted %v", alertsActive, tt.expectedAlertActive)
			}

			//Check if the data returned is as expected
			if !reflect.DeepEqual(data, tt.expectedData) {
				t.Errorf("ProcessAlerts() data = %v, wanted %v", data, tt.expectedData)
			}
		})
	}
}

func sortSliceByDate(data *[]models.HistoricalTradeAlertReturned) {
	sort.Slice(*data, func(i, j int) bool {
		return (*data)[i].Date.Before((*data)[j].Date)
	})
}
