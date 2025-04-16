package tests

import (
	"Price_Notification_System/models"
	"Price_Notification_System/service"
	"Price_Notification_System/store"
	"context"
	"errors"
	"testing"
	"time"
)

func TestTradeService_GetTradesByItem(t *testing.T) {

	type testDefGetTradesByItem struct {
		instanceName  string
		tradeStore    store.TradeStore
		item          string
		expectedData  []models.HistoricalTradeDataReturned
		expectedError error
		wantErr       bool
	}

	tests := []testDefGetTradesByItem{
		{instanceName: "test1 - Get trades for hulk item", tradeStore: store.NewInMemoryTradeStoreWithData(map[time.Time]models.HistoricalDataValues{
			time.Date(2024, 3, 10, 23, 20, 0, 0, time.UTC): {Object: "hulk", Price: 975},
			time.Date(2025, 1, 12, 3, 20, 0, 0, time.UTC):  {Object: "spider man", Price: 975},
			time.Date(2024, 5, 15, 3, 15, 30, 0, time.UTC): {Object: "hulk", Price: 854},
		}), item: "hulk",
			expectedData: []models.HistoricalTradeDataReturned{
				{Date: time.Date(2024, 3, 10, 23, 20, 0, 0, time.UTC), HistoricalDataValues: models.HistoricalDataValues{Object: "hulk", Price: 975}},
				{Date: time.Date(2024, 5, 15, 3, 15, 30, 0, time.UTC), HistoricalDataValues: models.HistoricalDataValues{Object: "hulk", Price: 854}},
			},
			expectedError: nil, wantErr: false},
		{instanceName: "test2 - Get trades for wolverine item", tradeStore: store.NewInMemoryTradeStoreWithData(map[time.Time]models.HistoricalDataValues{
			time.Date(2023, 3, 12, 23, 20, 0, 0, time.UTC): {Object: "wolverine", Price: 975},
			time.Date(2025, 1, 15, 3, 20, 0, 0, time.UTC):  {Object: "superman", Price: 250},
			time.Date(2024, 5, 18, 3, 15, 30, 0, time.UTC): {Object: "wolverine", Price: 790},
		}), item: "wolverine",
			expectedData: []models.HistoricalTradeDataReturned{
				{Date: time.Date(2023, 3, 12, 23, 20, 0, 0, time.UTC), HistoricalDataValues: models.HistoricalDataValues{Object: "wolverine", Price: 975}},
				{Date: time.Date(2024, 5, 18, 3, 15, 30, 0, time.UTC), HistoricalDataValues: models.HistoricalDataValues{Object: "wolverine", Price: 790}},
			},
			expectedError: nil, wantErr: false},
		{instanceName: "test3 - Ensure no matching trades are handled correctly", tradeStore: store.NewInMemoryTradeStoreWithData(map[time.Time]models.HistoricalDataValues{
			time.Date(2023, 3, 12, 23, 20, 0, 0, time.UTC): {Object: "wolverine", Price: 975},
			time.Date(2025, 1, 15, 3, 20, 0, 0, time.UTC):  {Object: "superman", Price: 250},
			time.Date(2024, 5, 18, 3, 15, 30, 0, time.UTC): {Object: "wolverine", Price: 790},
		}), item: "rob",
			expectedData:  []models.HistoricalTradeDataReturned{},
			expectedError: errors.New("no trades matching alerts found"), wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.instanceName, func(t *testing.T) {

			//Get the trades for the item
			ctx := context.Background()
			data, err := service.GetTradesByItem(ctx, tt.tradeStore, tt.item)

			//Check if the trades were returned correctly
			if !tt.wantErr && err != tt.expectedError {
				t.Errorf("GetTradesByItem() received error = %v, wanted %v", err, tt.expectedError)
			}
			if !tt.wantErr && len(data) != len(tt.expectedData) {
				t.Errorf("GetTradesByItem() received data = %v, wanted %v", data, tt.expectedData)
			}
		})
	}
}
