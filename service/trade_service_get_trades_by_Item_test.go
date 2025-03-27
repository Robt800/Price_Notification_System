package service

import (
	"Price_Notification_System/models"
	"Price_Notification_System/store"
	"context"
	"testing"
	"time"
)

func TestTradeService_GetTradesByItem(t *testing.T) {

	type testDefGetTradesByItem struct {
		instanceName   string
		mockTradeStore map[time.Time]models.HistoricalDataValues
		item           string
		expectedData   []models.HistoricalTradeDataReturned
		expectedError  error
		wantErr        bool
	}

	tests := []testDefGetTradesByItem{
		{instanceName: "test1 - Get trades for hulk item", mockTradeStore: map[time.Time]models.HistoricalDataValues{
			time.Date(2024, 1, 31, 10, 30, 20, 0, time.UTC): {Object: "hulk", Price: 975},
			time.Date(2025, 3, 23, 5, 45, 10, 0, time.UTC):  {Object: "he-man", Price: 1075},
			time.Date(2024, 5, 15, 3, 15, 30, 0, time.UTC):  {Object: "hulk", Price: 854},
		}, item: "hulk",
			expectedData: []models.HistoricalTradeDataReturned{
				{Date: time.Date(2024, 31, 10, 30, 20, 0, 0, time.UTC), HistoricalDataValues: models.HistoricalDataValues{Object: "hulk", Price: 975}},
				{Date: time.Date(2024, 5, 15, 3, 15, 30, 0, time.UTC), HistoricalDataValues: models.HistoricalDataValues{Object: "hulk", Price: 854}},
			},
			expectedError: nil, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.instanceName, func(t *testing.T) {

			//Create a new instance of the inMemoryTradeStore and set the data to the mockTradeStore
			testInMemoryTradeStore := store.NewInMemoryTradeStore().(*store.InMemoryTradeStore)
			testInMemoryTradeStore.TradeData = tt.mockTradeStore

			//Get the trades for the item
			ctx := context.Context(context.Background())
			data, err := GetTradesByItem(ctx, testInMemoryTradeStore, tt.item)

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
