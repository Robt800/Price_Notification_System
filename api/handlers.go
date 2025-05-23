package api

import (
	"Price_Notification_System/models"
	"Price_Notification_System/service"
	store "Price_Notification_System/store"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func GetTradesByItemHandler(ctx context.Context, itemTradeHistory store.TradeStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//Create key-value pairs (map) from the URL
		vars := mux.Vars(r)

		//Obtain from the url which item we are interested in reporting
		itemsToReport := vars["id"]

		//Create a var to store the results in
		var results []models.HistoricalTradeDataReturned
		var err error

		//Call the function from the service package
		results, err = service.GetTradesByItem(ctx, itemTradeHistory, itemsToReport)
		if err != nil {
			log.Println("The HTTP server failed to get the results due to error: ", err)
		}

		//Convert the results to JSON in readiness to respond with the results
		resultsJSON, err := json.Marshal(results)
		if err != nil {
			log.Println("The HTTP server failed to get the results due to error: ", err)
		}

		//Write the results to the
		_, errWrite := w.Write(resultsJSON)
		if errWrite != nil {
			log.Println("The HTTP server failed to get the results due to error: ", err)
		}
	}
}

func CreateNewAlertHandler(ctx context.Context, alertsDefined store.AlertDefStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Return a 200 OK response
		w.WriteHeader(http.StatusOK)

		//Create a var to store the response in
		var responseConfirmationString string
		var responseConfirmationStringJSON []byte

		//Create key-value pairs (map) from the URL
		vars := mux.Vars(r)

		//Obtain from the url which item we are interested in reporting - expected format is item,alertType,priceTrigger
		newAlertDetails := vars["id"]

		// Split out the new alert details
		item, alertType, priceTrigger, err := CreateNewAlertIDSplit(newAlertDetails)
		if err != nil {
			responseConfirmationString = fmt.Sprintf("There was an error trying to get the item details from the url.  Error: %v", err)
			responseConfirmationStringJSON, _ = json.Marshal(responseConfirmationString)
			_, err = w.Write(responseConfirmationStringJSON)
			if err != nil {
				log.Print("Error writing response: ", err)
			}
			return
		}

		//Call the function from the service package
		err = service.CreateNewAlert(ctx, alertsDefined, item, models.AlertValues{AlertType: alertType, PriceTrigger: priceTrigger})
		//results, err = service.ProcessAlerts(ctx, alertsDefined, itemsToReport)
		if err != nil {
			responseConfirmationString = fmt.Sprintf("The HTTP server failed to get the results due to error: %v", err)
			responseConfirmationStringJSON, _ = json.Marshal(responseConfirmationString)
			_, err = w.Write(responseConfirmationStringJSON)
			if err != nil {
				log.Print("Error writing response: ", err)
			}
			return
		}

		// If the alert was created successfully, return a confirmation message
		responseConfirmationString = fmt.Sprintf("The alert for item %s has been created successfully.\n The alert type is %d, and the price trigger is %d", item, alertType, priceTrigger)
		responseConfirmationStringJSON, _ = json.Marshal(responseConfirmationString)
		_, err = w.Write(responseConfirmationStringJSON)
		if err != nil {
			log.Print("Error writing response: ", err)
		}

	}
}

func GetAllDefinedAlertsHandler(ctx context.Context, alertsDefined store.AlertDefStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		dataAlertsDefined, err := alertsDefined.GetAllAlerts()
		if err != nil {
			marshalledError, _ := json.Marshal(err)
			_, errWrite := w.Write(marshalledError)
			if errWrite != nil {
				log.Println("The HTTP server failed to get the results due to error: ", err)
			}
			return
		}

		//Convert the results to JSON in readiness to respond with the results
		resultsJSON, err := json.Marshal(dataAlertsDefined)
		if err != nil {
			errJSON, _ := json.Marshal(err)
			_, _ = w.Write(errJSON)
			return
		}

		//Write the results to the
		_, errWrite := w.Write(resultsJSON)
		if errWrite != nil {
			errWriteJSON, _ := json.Marshal(errWrite)
			_, _ = w.Write(errWriteJSON)
		}
	}
}

func CreateNewAlertIDSplit(newAlertDetails string) (item string, alertType models.AlertType, priceTrigger int, err error) {
	// Create a temporary variable to hold the alert type as a string & price trigger as an int64
	var alertTypeAsString string
	var alertTypeInt int

	// Split the new alert details into item, alert type and price trigger
	alertDetails := strings.Split(newAlertDetails, ",")

	if len(alertDetails) != 3 {
		return "", 0, 0, fmt.Errorf("invalid alert details format")
	}

	item = alertDetails[0]
	alertTypeAsString = alertDetails[1]
	priceTrigger, err = strconv.Atoi(alertDetails[2])
	if err != nil {
		return "", 0, 0, fmt.Errorf("invalid new alert details format")
	}

	// Convert the alert type string to the AlertType enum
	alertTypeInt, err = strconv.Atoi(alertTypeAsString)
	if err != nil {
		return "", 0, 0, fmt.Errorf("invalid alert type format")
	}

	alertType = models.AlertType(alertTypeInt)

	return item, alertType, priceTrigger, nil
}
