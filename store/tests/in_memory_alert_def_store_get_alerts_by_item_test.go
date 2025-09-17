package tests

import (
	"Price_Notification_System/models"
	"Price_Notification_System/store"
	"testing"
)

func TestInMemoryAlertStore_GetAlertsByItem(t *testing.T) {
	type testDefGetAlertsByItem struct {
		instanceName    string
		alertDefStore   store.AlertDefStore
		itemToGetAlerts string
		expected        []models.AlertsByItemReturned
		wantErr         bool
	}

	tests := []testDefGetAlertsByItem{
		{instanceName: "test1 - Get alerts for hulk item",
			alertDefStore: store.NewInMemoryAlertStoreWithData(&[]models.AlertDef{
				{Item: "hulk", AlertValues: models.AlertValues{AlertType: models.PriceAlertLowPrice, PriceTrigger: 975}, EmailRecipient: "rob@test.com"},
				{Item: "wolverine", AlertValues: models.AlertValues{AlertType: models.PriceAlertLowPrice, PriceTrigger: 600}, EmailRecipient: "sarah@test.com"},
				{Item: "hulk", AlertValues: models.AlertValues{AlertType: models.PriceAlertLowPrice, PriceTrigger: 600}, EmailRecipient: "chris@test.com"},
			}),
			itemToGetAlerts: "hulk",
			expected: []models.AlertsByItemReturned{
				{Item: "hulk", AlertValues: models.AlertValues{AlertType: models.PriceAlertLowPrice, PriceTrigger: 975}, EmailRecipient: "rob@test.com"},
				{Item: "hulk", AlertValues: models.AlertValues{AlertType: models.PriceAlertLowPrice, PriceTrigger: 600}, EmailRecipient: "chris@test.com"},
			},
			wantErr: false,
		},

		{instanceName: "test2 - Get alerts for wolverine item",
			alertDefStore: store.NewInMemoryAlertStoreWithData(&[]models.AlertDef{
				{Item: "hulk", AlertValues: models.AlertValues{AlertType: models.PriceAlertLowPrice, PriceTrigger: 975}, EmailRecipient: "jayne@test.com"},
				{Item: "wolverine", AlertValues: models.AlertValues{AlertType: models.PriceAlertLowPrice, PriceTrigger: 600}, EmailRecipient: "ted@test.com"},
				{Item: "wolverine", AlertValues: models.AlertValues{AlertType: models.PriceAlertHighPrice, PriceTrigger: 1200}, EmailRecipient: "merv@test.com"},
				{Item: "wolverine", AlertValues: models.AlertValues{AlertType: models.PriceAlertLowPrice, PriceTrigger: 200}, EmailRecipient: "shirley@test.com"},
			}),
			itemToGetAlerts: "wolverine",
			expected: []models.AlertsByItemReturned{
				{Item: "wolverine", AlertValues: models.AlertValues{AlertType: models.PriceAlertLowPrice, PriceTrigger: 600}, EmailRecipient: "ted@test.com"},
				{Item: "wolverine", AlertValues: models.AlertValues{AlertType: models.PriceAlertHighPrice, PriceTrigger: 1200}, EmailRecipient: "merv@test.com"},
				{Item: "wolverine", AlertValues: models.AlertValues{AlertType: models.PriceAlertLowPrice, PriceTrigger: 200}, EmailRecipient: "shirley@test.com"},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.instanceName, func(t *testing.T) {
			got, err := tt.alertDefStore.GetAlertsByItem(tt.itemToGetAlerts)
			if (err != nil) != tt.wantErr {
				t.Errorf("InMemoryAlertStore.GetAlertsByItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(got) != len(tt.expected) {
				t.Errorf("InMemoryAlertStore.GetAlertsByItem() = %v, expected %v", got, tt.expected)
				return
			}

			for i := range got {
				if got[i].Item != tt.expected[i].Item || got[i].AlertValues != tt.expected[i].AlertValues {
					t.Errorf("InMemoryAlertStore.GetAlertsByItem() = %v, expected %v", got, tt.expected)
					return
				}
			}
		})
	}

}
