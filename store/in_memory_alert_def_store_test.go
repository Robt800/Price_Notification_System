package store

import (
	"Price_Notification_System/models"
	"reflect"
	"testing"
)

func TestInMemoryAlertDefStore_AddAlert(t *testing.T) {

	type testDefAddAlert struct {
		instanceName    string
		mockStoreAlerts map[string]models.AlertValues
		itemToAlert     string
		newAlertDef     models.AlertValues
		expected        map[string]models.AlertValues
		wantErr         bool
	}

	tests := []testDefAddAlert{
		{instanceName: "test1 - Add 1st alert to empty store", mockStoreAlerts: map[string]models.AlertValues{}, itemToAlert: "test item1",
			newAlertDef: models.AlertValues{AlertType: models.PriceAlertLowPrice, PriceTrigger: 975},
			expected:    map[string]models.AlertValues{"test item1": {AlertType: models.PriceAlertLowPrice, PriceTrigger: 975}}, wantErr: false},

		{instanceName: "test2 - Add alert to existing store", mockStoreAlerts: map[string]models.AlertValues{
			"test item1": {AlertType: models.PriceAlertLowPrice, PriceTrigger: 975},
		}, itemToAlert: "test item2",
			newAlertDef: models.AlertValues{AlertType: models.PriceAlertHighPrice, PriceTrigger: 1075},
			expected: map[string]models.AlertValues{"test item1": {AlertType: models.PriceAlertLowPrice, PriceTrigger: 975},
				"test item2": {AlertType: models.PriceAlertHighPrice, PriceTrigger: 1075}},
			wantErr: false},

		{instanceName: "test3 - Add a duplicate alert to existing store", mockStoreAlerts: map[string]models.AlertValues{
			"test item3": {AlertType: models.PriceAlertLowPrice, PriceTrigger: 975},
		}, itemToAlert: "test item3",
			newAlertDef: models.AlertValues{AlertType: models.PriceAlertLowPrice, PriceTrigger: 975},
			expected:    map[string]models.AlertValues{"test item3": {AlertType: models.PriceAlertLowPrice, PriceTrigger: 975}},
			wantErr:     false},
	}

	for _, tt := range tests {
		t.Run(tt.instanceName, func(t *testing.T) {

			//Create a new instance of the inMemoryAlertStore and set the data to the mockStoreAlerts
			testInMemoryAlertStore := NewInMemoryAlertStore().(*InMemoryAlertStore)
			testInMemoryAlertStore.data = tt.mockStoreAlerts

			//Add the alert to the inMemoryAlertStore
			testInMemoryAlertStore.AddAlert(tt.itemToAlert, tt.newAlertDef)

			//Check if the alert was added to the inMemoryAlertStore correctly
			if !reflect.DeepEqual(testInMemoryAlertStore.data, tt.expected) {
				t.Errorf("AddAlert() received = %v\n, wanted %v", testInMemoryAlertStore.data, tt.expected)
			}
		})
	}
}
