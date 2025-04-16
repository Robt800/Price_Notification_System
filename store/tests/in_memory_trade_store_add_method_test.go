package tests

import (
	"Price_Notification_System/models"
	"Price_Notification_System/store"
	"reflect"
	"testing"
	"time"
)

func TestInMemoryTradeStoreAddTrade(t *testing.T) {

	// testDefAddMethod is a struct that contains the parameters for the test
	type testDefAddMethod struct {
		instanceName string
		tradeStore   store.TradeStore
		tradeTime    time.Time
		tradeValues  models.HistoricalDataValues
		expected     store.TradeStore
		wantErr      bool
	}

	// testCases is a slice of testDef structs that contains the parameters for each test
	testCasesAddMethod := []testDefAddMethod{
		{instanceName: "Add 1st trade",
			tradeStore:  store.NewInMemoryTradeStoreWithData(&map[time.Time]models.HistoricalDataValues{}),
			tradeTime:   time.Date(2025, time.January, 20, 19, 34, 0, 0, time.UTC),
			tradeValues: models.HistoricalDataValues{Object: "testObject1", Price: 900},
			expected: store.NewInMemoryTradeStoreWithData(&map[time.Time]models.HistoricalDataValues{
				time.Date(2025, time.January, 20, 19, 34, 0, 0, time.UTC): {Object: "testObject1", Price: 900},
			}),
			wantErr: false},

		{instanceName: "Add 2nd trade",
			tradeStore: store.NewInMemoryTradeStoreWithData(&map[time.Time]models.HistoricalDataValues{
				time.Date(2025, time.January, 20, 19, 34, 0, 0, time.UTC): {Object: "testObject1", Price: 900},
			}),
			tradeTime:   time.Date(2024, time.February, 27, 8, 38, 0, 0, time.UTC),
			tradeValues: models.HistoricalDataValues{Object: "testObject2", Price: 600},
			expected: store.NewInMemoryTradeStoreWithData(&map[time.Time]models.HistoricalDataValues{
				time.Date(2025, time.January, 20, 19, 34, 0, 0, time.UTC): {Object: "testObject1", Price: 900},
				time.Date(2024, time.February, 27, 8, 38, 0, 0, time.UTC): {Object: "testObject2", Price: 600},
			}),
			wantErr: false},

		{instanceName: "Add 3rd trade",
			tradeStore: store.NewInMemoryTradeStoreWithData(&map[time.Time]models.HistoricalDataValues{
				time.Date(2026, time.December, 20, 14, 39, 40, 0, time.UTC): {Object: "testObject1", Price: 300},
				time.Date(2025, time.March, 05, 11, 15, 40, 0, time.UTC):    {Object: "testObject1", Price: 500},
			}),
			tradeTime:   time.Date(2023, time.September, 14, 6, 15, 30, 0, time.UTC),
			tradeValues: models.HistoricalDataValues{Object: "testObject3", Price: 1000},
			expected: store.NewInMemoryTradeStoreWithData(&map[time.Time]models.HistoricalDataValues{
				time.Date(2026, time.December, 20, 14, 39, 40, 0, time.UTC): {Object: "testObject1", Price: 300},
				time.Date(2025, time.March, 05, 11, 15, 40, 0, time.UTC):    {Object: "testObject1", Price: 500},
				time.Date(2023, time.September, 14, 6, 15, 30, 0, time.UTC): {Object: "testObject3", Price: 1000},
			}),
			wantErr: false},
	}

	//Iterate through the test cases regarding the AddTrade method
	for _, tt := range testCasesAddMethod {
		t.Run(tt.instanceName, func(t *testing.T) {

			// add the trade to the inMemoryTradeStore
			tt.tradeStore.AddTrade(tt.tradeTime, tt.tradeValues)

			// check if the trade was added to the inMemoryTradeStore
			if !reflect.DeepEqual(tt.tradeStore, tt.expected) {
				t.Errorf("AddTrade() = %v, want %v", tt.tradeStore, tt.expected)
			}
		})
	}

}
