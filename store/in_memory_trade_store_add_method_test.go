package store

import (
	"Price_Notification_System/models"
	"reflect"
	"testing"
	"time"
)

func TestInMemoryTradeStoreAddTrade(t *testing.T) {

	// testDefAddMethod is a struct that contains the parameters for the test
	type testDefAddMethod struct {
		instanceName string
		tradeTime    time.Time
		tradeValues  models.HistoricalDataValues
		expected     map[time.Time]models.HistoricalDataValues
		wantErr      bool
	}

	// testCases is a slice of testDef structs that contains the parameters for each test
	testCasesAddMethod := []testDefAddMethod{
		{instanceName: "Add 1st trade", tradeTime: time.Date(2025, time.January, 20, 19, 34, 0, 0, time.UTC),
			tradeValues: models.HistoricalDataValues{Object: "testObject1", Price: 900},
			expected: map[time.Time]models.HistoricalDataValues{
				time.Date(2025, time.January, 20, 19, 34, 0, 0, time.UTC): {Object: "testObject1", Price: 900}},
			wantErr: false},
		{instanceName: "Add 2nd trade", tradeTime: time.Date(2024, time.February, 27, 8, 38, 0, 0, time.UTC),
			tradeValues: models.HistoricalDataValues{Object: "testObject2", Price: 600},
			expected: map[time.Time]models.HistoricalDataValues{
				time.Date(2024, time.February, 27, 8, 38, 0, 0, time.UTC): {Object: "testObject2", Price: 600}},
			wantErr: false},
		{instanceName: "Add 2nd trade", tradeTime: time.Date(2023, time.September, 14, 6, 15, 30, 0, time.UTC),
			tradeValues: models.HistoricalDataValues{Object: "testObject3", Price: 1000},
			expected: map[time.Time]models.HistoricalDataValues{
				time.Date(2023, time.September, 14, 6, 15, 30, 0, time.UTC): {Object: "testObject3", Price: 1000}},
			wantErr: false},
	}

	//Iterate through the test cases regarding the AddTrade method
	for _, tt := range testCasesAddMethod {
		t.Run(tt.instanceName, func(t *testing.T) {

			// create a new instance of the inMemoryTradeStore
			testInMemoryTradeStore := NewInMemoryTradeStore().(*inMemoryTradeStore)

			// add the trade to the inMemoryTradeStore
			testInMemoryTradeStore.AddTrade(tt.tradeTime, tt.tradeValues)

			// check if the trade was added to the inMemoryTradeStore
			if !reflect.DeepEqual(testInMemoryTradeStore.tradeData, tt.expected) {
				t.Errorf("AddTrade() = %v, want %v", testInMemoryTradeStore.tradeData, tt.expected)
			}
		})
	}

}
