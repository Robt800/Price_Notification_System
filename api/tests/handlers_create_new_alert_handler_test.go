package api

import (
	"Price_Notification_System/api"
	"Price_Notification_System/models"
	"Price_Notification_System/store"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestCreateNewAlertHandler(t *testing.T) {

	type testDef struct {
		testInstanceName      string
		alertsDefined         store.AlertDefStore
		url                   string
		body                  url.Values
		expectedAlertsDefined store.AlertDefStore
		expectedResponse      string
		wantErr               bool
	}

	// Define a slice of testDef structs
	tests := []testDef{
		{
			testInstanceName: "test1 - Create new alert",
			alertsDefined:    store.NewInMemoryAlertStore(),
			url:              "/items/Batman figure/alerts",
			body: url.Values{
				"alertType":    []string{"0"},
				"priceTrigger": []string{"100"},
			},
			expectedAlertsDefined: store.NewInMemoryAlertStoreWithData(&[]models.AlertDef{
				{Item: "Batman figure", AlertValues: models.AlertValues{AlertType: 0, PriceTrigger: 100}},
			}),
			expectedResponse: fmt.Sprintf("The alert for item Batman figure has been created successfully.\n The alert type is %d, and the price trigger is %d", 0, 100),
			wantErr:          false,
		},
	}

	// Iterate over the test cases
	for _, tt := range tests {
		t.Run(tt.testInstanceName, func(t *testing.T) {

			// Create a new HTTP request to pass to the handler
			req, err := http.NewRequest("POST", tt.url, strings.NewReader(tt.body.Encode()))
			if err != nil {
				t.Fatal(err)
			}

			// Set the content type to application/x-www-form-urlencoded
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			// Create a new ResponseRecorder to record the response
			rr := httptest.NewRecorder()

			// Call the handler
			handler := api.CreateNewAlertHandler(context.Background(), tt.alertsDefined)

			// Serve the HTTP request
			handler.ServeHTTP(rr, req)

			// Check the response status code
			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}

			// Check the response body

			// Prepare expected response as JSON
			expectedResponseJSON, err := json.Marshal(tt.expectedResponse)
			if err != nil {
				t.Fatalf("could not marshal expected response: %v", err)
			}
			if rr.Body.String() != string(expectedResponseJSON) {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), string(expectedResponseJSON))
			}

			// Check if the alert was added to the store correctly
			// Get the alerts from the store - actual
			gotAlerts, err := tt.alertsDefined.GetAllAlerts()
			if err != nil {
				if !tt.wantErr {
					t.Errorf("unexpected error: %v", err)
				}
				return
			}
			if tt.wantErr {
				t.Errorf("expected error but got none")
			}

			// Check if the alert was added to the store correctly
			// Get the alerts from the store - expected
			expectedAlerts, err := tt.expectedAlertsDefined.GetAllAlerts()
			if err != nil {
				if !tt.wantErr {
					t.Errorf("unexpected error: %v", err)
				}
				return
			}
			if tt.wantErr {
				t.Errorf("expected error but got none")
			}

			if !reflect.DeepEqual(gotAlerts, expectedAlerts) {
				t.Errorf("alertsDefined store does not match expected: got %v want %v",
					tt.alertsDefined, tt.expectedAlertsDefined)
			}
		})
	}

}
