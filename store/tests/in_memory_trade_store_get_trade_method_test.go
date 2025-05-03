package tests

import (
	"Price_Notification_System/models"
	"Price_Notification_System/store"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestInMemoryTradeStore_GetTradeByItem(t *testing.T) {

	// testDefGetTradeMethod is a struct that contains the parameters for the test
	type testDefGetTradeMethod struct {
		instanceName string
		dataStore    store.TradeStore
		itemRequired string
		expected     []models.HistoricalTradeDataReturned
		wantErr      bool
	}

	// testCases is a slice of testDef structs that contains the parameters for each test
	testCasesGetTradeMethod := []testDefGetTradeMethod{
		{instanceName: "Get trade by item1",
			dataStore: store.NewInMemoryTradeStoreWithData(&map[time.Time]models.HistoricalDataValues{
				time.Date(2025, time.January, 20, 19, 34, 0, 0, time.UTC):   {Object: "testObject1", Price: 900},
				time.Date(2024, time.February, 27, 8, 38, 0, 0, time.UTC):   {Object: "testObject2", Price: 600},
				time.Date(2023, time.September, 14, 6, 15, 30, 0, time.UTC): {Object: "testObject1", Price: 1000},
				time.Date(2025, time.March, 14, 15, 15, 30, 0, time.UTC):    {Object: "testObject1", Price: 800},
				time.Date(2024, time.October, 28, 13, 30, 45, 0, time.UTC):  {Object: "testObject3", Price: 750},
				time.Date(2023, time.April, 5, 22, 30, 45, 0, time.UTC):     {Object: "testObject1", Price: 975},
			}),
			itemRequired: "testObject1",
			expected: []models.HistoricalTradeDataReturned{
				{Date: time.Date(2025, time.January, 20, 19, 34, 0, 0, time.UTC), HistoricalDataValues: models.HistoricalDataValues{Object: "testObject1", Price: 900}},
				{Date: time.Date(2023, time.September, 14, 6, 15, 30, 0, time.UTC), HistoricalDataValues: models.HistoricalDataValues{Object: "testObject1", Price: 1000}},
				{Date: time.Date(2025, time.March, 14, 15, 15, 30, 0, time.UTC), HistoricalDataValues: models.HistoricalDataValues{Object: "testObject1", Price: 800}},
				{Date: time.Date(2023, time.April, 5, 22, 30, 45, 0, time.UTC), HistoricalDataValues: models.HistoricalDataValues{Object: "testObject1", Price: 975}},
			},
			wantErr: false},

		{instanceName: "Get trade by item2",
			dataStore: store.NewInMemoryTradeStoreWithData(&map[time.Time]models.HistoricalDataValues{
				time.Date(2025, time.January, 20, 19, 34, 0, 0, time.UTC):   {Object: "Hulk", Price: 900},
				time.Date(2024, time.February, 10, 8, 38, 0, 0, time.UTC):   {Object: "Thor", Price: 600},
				time.Date(2023, time.September, 11, 6, 15, 30, 0, time.UTC): {Object: "IronMan", Price: 1000},
				time.Date(2025, time.March, 24, 15, 15, 30, 0, time.UTC):    {Object: "Hulk", Price: 800},
			}),
			itemRequired: "Hulk",
			expected: []models.HistoricalTradeDataReturned{
				{Date: time.Date(2025, time.January, 20, 19, 34, 0, 0, time.UTC), HistoricalDataValues: models.HistoricalDataValues{Object: "Hulk", Price: 900}},
				{Date: time.Date(2025, time.March, 24, 15, 15, 30, 0, time.UTC), HistoricalDataValues: models.HistoricalDataValues{Object: "Hulk", Price: 800}},
			},
			wantErr: false},
	}

	//Iterate through the test cases regarding the GetTradeByItem method
	for _, tt := range testCasesGetTradeMethod {
		t.Run(tt.instanceName, func(t *testing.T) {

			// get the trade by item from the inMemoryTradeStore
			data, err := tt.dataStore.GetTradeByItem(tt.itemRequired)

			//Sort the data
			sortSlice(data)
			sortSlice(tt.expected)

			// check if the trade was retrieved from the inMemoryTradeStore
			if !reflect.DeepEqual(data, tt.expected) {
				t.Errorf("GetTradeByItem() = %v,\n want %v", data, tt.expected)
			}

			// check if the error returned by the GetTradeByItem method is as expected
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTradeByItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func sortSlice(data []models.HistoricalTradeDataReturned) []models.HistoricalTradeDataReturned {
	sort.Slice(data, func(i, j int) bool {
		return data[i].Date.Before(data[j].Date)
	})
	return data
}
