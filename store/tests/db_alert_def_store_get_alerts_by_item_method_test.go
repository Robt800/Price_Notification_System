package tests

import (
	"Price_Notification_System/config"
	"Price_Notification_System/models"
	"Price_Notification_System/store"
	"reflect"
	"testing"
)

func TestGetAlertsByItem(t *testing.T) {

	//Load the environment variables
	enVariables, err := config.LoadEnvVariables()
	if err != nil {
		t.Fatalf("Error loading environment variables: %v", err)
	}
	//Load the DB connection string
	dbConnStr := config.LoadDBConnectionStr(enVariables)

	alertDefStore, errorCreatingStore := store.NewDBAlertStore(dbConnStr)

	type testDefGetAlertsByItem struct {
		instanceName       string
		item               string
		alertDefStore      store.AlertDefStore
		errorCreatingStore error
		expected           []models.AlertsByItemReturned
		wantErr            bool
	}

	tests := []testDefGetAlertsByItem{
		{
			instanceName:       "test1 - Get alerts for item with no alerts",
			item:               "Sandman",
			alertDefStore:      alertDefStore,
			errorCreatingStore: errorCreatingStore,
			expected:           []models.AlertsByItemReturned{},
			wantErr:            false,
		},
		{
			instanceName:       "test2 - Get alerts for item with seven alerts",
			item:               "Batman",
			alertDefStore:      alertDefStore,
			errorCreatingStore: errorCreatingStore,
			expected: []models.AlertsByItemReturned{
				{Item: "Batman", AlertValues: models.AlertValues{AlertType: models.PriceAlertLowPrice, PriceTrigger: 500}},
				{Item: "Batman", AlertValues: models.AlertValues{AlertType: models.PriceAlertLowPrice, PriceTrigger: 500}},
				{Item: "Batman", AlertValues: models.AlertValues{AlertType: models.PriceAlertLowPrice, PriceTrigger: 500}},
				{Item: "Batman", AlertValues: models.AlertValues{AlertType: models.PriceAlertLowPrice, PriceTrigger: 500}},
				{Item: "Batman", AlertValues: models.AlertValues{AlertType: models.PriceAlertLowPrice, PriceTrigger: 500}},
				{Item: "Batman", AlertValues: models.AlertValues{AlertType: models.PriceAlertLowPrice, PriceTrigger: 500}},
				{Item: "Batman", AlertValues: models.AlertValues{AlertType: models.PriceAlertLowPrice, PriceTrigger: 200}},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.instanceName, func(t *testing.T) {

			//Get the data from DBAlertStore
			data, errGettingAlerts := tt.alertDefStore.GetAlertsByItem(tt.item)
			if (errGettingAlerts != nil) != tt.wantErr {
				t.Errorf("GetAlertByItem() error = %v, wantErr %v", errGettingAlerts, tt.wantErr)
			}

			// check if the trade was retrieved from the inMemoryTradeStore
			if !reflect.DeepEqual(data, tt.expected) {
				t.Errorf("GetAlertsByItem() = %v,\n want %v", data, tt.expected)
			}

		})
	}

}
