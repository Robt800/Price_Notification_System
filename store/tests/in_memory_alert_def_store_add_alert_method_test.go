package tests

import (
	"Price_Notification_System/models"
	"Price_Notification_System/store"
	"reflect"
	"testing"
)

func TestInMemoryAlertDefStore_AddAlert(t *testing.T) {

	type testDefAddAlert struct {
		instanceName   string
		alertDefStore  store.AlertDefStore
		itemToAlert    string
		newAlertDef    models.AlertValues
		emailRecipient string
		expected       store.AlertDefStore
		wantErr        bool
	}

	tests := []testDefAddAlert{
		{instanceName: "test1 - Add 1st alert to empty store",
			alertDefStore:  store.NewInMemoryAlertStoreWithData(&[]models.AlertDef{}),
			itemToAlert:    "test item1",
			newAlertDef:    models.AlertValues{AlertType: models.PriceAlertLowPrice, PriceTrigger: 975},
			emailRecipient: "rob@test.com",
			expected: store.NewInMemoryAlertStoreWithData(&[]models.AlertDef{
				{Item: "test item1",
					AlertValues:    models.AlertValues{AlertType: models.PriceAlertLowPrice, PriceTrigger: 975},
					EmailRecipient: "rob@test.com"},
			}),
			wantErr: false},

		{instanceName: "test2 - Add alert to existing store",
			alertDefStore: store.NewInMemoryAlertStoreWithData(&[]models.AlertDef{
				{Item: "test item1",
					AlertValues: models.AlertValues{
						AlertType:    models.PriceAlertLowPrice,
						PriceTrigger: 975},
					EmailRecipient: "rob@test.com"},
			}),
			itemToAlert: "test item2",
			newAlertDef: models.AlertValues{
				AlertType:    models.PriceAlertHighPrice,
				PriceTrigger: 1075},
			emailRecipient: "ted@test.com",
			expected: store.NewInMemoryAlertStoreWithData(&[]models.AlertDef{
				{Item: "test item1",
					AlertValues: models.AlertValues{
						AlertType:    models.PriceAlertLowPrice,
						PriceTrigger: 975},
					EmailRecipient: "rob@test.com"},
				{Item: "test item2",
					AlertValues: models.AlertValues{
						AlertType:    models.PriceAlertHighPrice,
						PriceTrigger: 1075},
					EmailRecipient: "ted@test.com"},
			}),
			wantErr: false},

		{instanceName: "test3 - Add a duplicate alert to existing store",
			alertDefStore: store.NewInMemoryAlertStoreWithData(&[]models.AlertDef{
				{Item: "test item1", AlertValues: models.AlertValues{AlertType: models.PriceAlertLowPrice, PriceTrigger: 975}, EmailRecipient: "ted@test.com"},
				{Item: "test item2", AlertValues: models.AlertValues{AlertType: models.PriceAlertHighPrice, PriceTrigger: 1075}, EmailRecipient: "rob@test.com"},
				{Item: "test item3", AlertValues: models.AlertValues{AlertType: models.PriceAlertHighPrice, PriceTrigger: 1100}, EmailRecipient: "sarah@test.com"},
			}),
			itemToAlert:    "test item3",
			newAlertDef:    models.AlertValues{AlertType: models.PriceAlertHighPrice, PriceTrigger: 1100},
			emailRecipient: "sarah@test.com",
			expected: store.NewInMemoryAlertStoreWithData(&[]models.AlertDef{
				{Item: "test item1", AlertValues: models.AlertValues{AlertType: models.PriceAlertLowPrice, PriceTrigger: 975}, EmailRecipient: "ted@test.com"},
				{Item: "test item2", AlertValues: models.AlertValues{AlertType: models.PriceAlertHighPrice, PriceTrigger: 1075}, EmailRecipient: "rob@test.com"},
				{Item: "test item3", AlertValues: models.AlertValues{AlertType: models.PriceAlertHighPrice, PriceTrigger: 1100}, EmailRecipient: "sarah@test.com"},
				{Item: "test item3", AlertValues: models.AlertValues{AlertType: models.PriceAlertHighPrice, PriceTrigger: 1100}, EmailRecipient: "sarah@test.com"},
			}),
			wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.instanceName, func(t *testing.T) {

			//Add the alert to the inMemoryAlertStore
			tt.alertDefStore.AddAlert(tt.itemToAlert, tt.newAlertDef, tt.emailRecipient)

			//Check if the alert was added to the inMemoryAlertStore correctly
			if !reflect.DeepEqual(tt.alertDefStore, tt.expected) {
				t.Errorf("AddAlert() received = %v\n, wanted %v", tt.alertDefStore, tt.expected)
			}
		})
	}
}
