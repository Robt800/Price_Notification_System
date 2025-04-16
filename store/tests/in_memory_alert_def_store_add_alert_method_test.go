package tests

import (
	"Price_Notification_System/models"
	"Price_Notification_System/store"
	"reflect"
	"testing"
)

func TestInMemoryAlertDefStore_AddAlert(t *testing.T) {

	type testDefAddAlert struct {
		instanceName  string
		alertDefStore store.AlertDefStore
		itemToAlert   string
		newAlertDef   models.AlertValues
		expected      store.AlertDefStore
		wantErr       bool
	}

	tests := []testDefAddAlert{
		{instanceName: "test1 - Add 1st alert to empty store",
			alertDefStore: store.NewInMemoryAlertStoreWithData(&[]models.AlertDef{}),
			itemToAlert:   "test item1",
			newAlertDef:   models.AlertValues{AlertType: models.PriceAlertLowPrice, PriceTrigger: 975},
			expected: store.NewInMemoryAlertStoreWithData(&[]models.AlertDef{
				{Item: "test item1", AlertValues: models.AlertValues{AlertType: models.PriceAlertLowPrice, PriceTrigger: 975}},
			}),
			wantErr: false},

		{instanceName: "test2 - Add alert to existing store",
			alertDefStore: store.NewInMemoryAlertStoreWithData(&[]models.AlertDef{
				{Item: "test item1", AlertValues: models.AlertValues{AlertType: models.PriceAlertLowPrice, PriceTrigger: 975}},
			}),
			itemToAlert: "test item2",
			newAlertDef: models.AlertValues{AlertType: models.PriceAlertHighPrice, PriceTrigger: 1075},
			expected: store.NewInMemoryAlertStoreWithData(&[]models.AlertDef{
				{Item: "test item1", AlertValues: models.AlertValues{AlertType: models.PriceAlertLowPrice, PriceTrigger: 975}},
				{Item: "test item2", AlertValues: models.AlertValues{AlertType: models.PriceAlertHighPrice, PriceTrigger: 1075}},
			}),
			wantErr: false},

		{instanceName: "test3 - Add a duplicate alert to existing store",
			alertDefStore: store.NewInMemoryAlertStoreWithData(&[]models.AlertDef{
				{Item: "test item1", AlertValues: models.AlertValues{AlertType: models.PriceAlertLowPrice, PriceTrigger: 975}},
				{Item: "test item2", AlertValues: models.AlertValues{AlertType: models.PriceAlertHighPrice, PriceTrigger: 1075}},
				{Item: "test item3", AlertValues: models.AlertValues{AlertType: models.PriceAlertHighPrice, PriceTrigger: 1100}},
			}),
			itemToAlert: "test item3",
			newAlertDef: models.AlertValues{AlertType: models.PriceAlertHighPrice, PriceTrigger: 1100},
			expected: store.NewInMemoryAlertStoreWithData(&[]models.AlertDef{
				{Item: "test item1", AlertValues: models.AlertValues{AlertType: models.PriceAlertLowPrice, PriceTrigger: 975}},
				{Item: "test item2", AlertValues: models.AlertValues{AlertType: models.PriceAlertHighPrice, PriceTrigger: 1075}},
				{Item: "test item3", AlertValues: models.AlertValues{AlertType: models.PriceAlertHighPrice, PriceTrigger: 1100}},
				{Item: "test item3", AlertValues: models.AlertValues{AlertType: models.PriceAlertHighPrice, PriceTrigger: 1100}},
			}),
			wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.instanceName, func(t *testing.T) {

			//Add the alert to the inMemoryAlertStore
			tt.alertDefStore.AddAlert(tt.itemToAlert, tt.newAlertDef)

			//Check if the alert was added to the inMemoryAlertStore correctly
			if !reflect.DeepEqual(tt.alertDefStore, tt.expected) {
				t.Errorf("AddAlert() received = %v\n, wanted %v", tt.alertDefStore, tt.expected)
			}
		})
	}
}
