package api

import (
	"Price_Notification_System/api"
	"Price_Notification_System/models"
	"Price_Notification_System/store"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllDefinedAlertsHandler(t *testing.T) {

	type testDef struct {
		testInstanceName string
		alertsDefined    store.AlertDefStore
		expected         []byte
		wantErr          bool
	}

	// Define a slice of testDef structs
	tests := []testDef{
		{testInstanceName: "test1 - Get all defined alerts",
			alertsDefined: store.NewInMemoryAlertStoreWithData(&[]models.AlertDef{
				{Item: "item1", AlertValues: models.AlertValues{AlertType: models.PriceAlertLowPrice, PriceTrigger: 100}},
				{Item: "item2", AlertValues: models.AlertValues{AlertType: models.PriceAlertHighPrice, PriceTrigger: 200}},
			}),
			expected: []byte(`{"data":[{"Item":"item1","AlertType":0,"PriceTrigger":100},{"Item":"item2","AlertType":1,"PriceTrigger":200}]}`),
			wantErr:  false,
		},
		{testInstanceName: "test2 - Get all defined alerts",
			alertsDefined: store.NewInMemoryAlertStoreWithData(&[]models.AlertDef{
				{Item: "item3", AlertValues: models.AlertValues{AlertType: 0, PriceTrigger: 300}},
				{Item: "item4", AlertValues: models.AlertValues{AlertType: 1, PriceTrigger: 400}},
			}),
			expected: []byte(`{"data":[{"Item":"item3","AlertType":0,"PriceTrigger":300},{"Item":"item4","AlertType":1,"PriceTrigger":400}]}`),
			wantErr:  false,
		},
	}

	// Iterate over the test cases
	for _, tt := range tests {
		t.Run(tt.testInstanceName, func(t *testing.T) {

			// Create a new HTTP request to pass to the handler
			req, err := http.NewRequest("GET", "/all_defined_alerts", nil)
			if err != nil {
				t.Fatal(err)
			}

			// Create a new ResponseRecorder to record the response
			rr := httptest.NewRecorder()

			// Call the handler
			handler := api.GetAllDefinedAlertsHandler(context.Background(), tt.alertsDefined)

			// Serve the HTTP request
			handler.ServeHTTP(rr, req)

			// Check the response status code
			if rr.Code != http.StatusOK {
				t.Errorf("expected status 200 OK, got %v", rr.Code)
			}

			// Check the response body
			got := rr.Body.Bytes()
			if string(got) != string(tt.expected) {
				t.Errorf("expected response body %s, got %s", string(tt.expected), string(got))
			}

		})
	}

}
