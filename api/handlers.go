package api

import (
	"Price_Notification_System/models"
	"Price_Notification_System/service"
	store "Price_Notification_System/store"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func GetTradesByItemHandler(ctx context.Context, itemTradeHistory store.TradeStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//Create key-value pairs (map) from the URL
		vars := mux.Vars(r)

		//Obtain from the url which item we are interested in reporting
		itemsToReport := vars["item"]

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

		//Create a var to store the response in
		var responseConfirmationString string
		var responseConfirmationStringJSON []byte

		//Obtain from the url which item we are interested in reporting - expected format is item,alertType,priceTrigger
		newAlertItem, err := CreateNewAlertObtainItemFromURL(r)
		if err != nil {
			responseConfirmationString = fmt.Sprintf("There was an error trying to get the item details from the HTTP request.  Error: %v", err)
			responseConfirmationStringJSON, _ = json.Marshal(responseConfirmationString)
			_, err = w.Write(responseConfirmationStringJSON)
			if err != nil {
				log.Print("Error writing response: ", err)
			}
			return
		}

		// Split out the new alert details - this is expected to be in the request body as key value pairs (alertType, priceTrigger) - of the format url encoded string
		alertType, priceTrigger, emailRecipient, err := CreateNewAlertRequestBodyParameters(r)
		if err != nil {
			responseConfirmationString = fmt.Sprintf("There was an error trying to get the item details from the HTTP request.  Error: %v", err)
			responseConfirmationStringJSON, _ = json.Marshal(responseConfirmationString)
			_, err = w.Write(responseConfirmationStringJSON)
			if err != nil {
				log.Print("Error writing response: ", err)
			}
			return
		}

		//Call the function from the service package
		err = service.CreateNewAlert(ctx, alertsDefined, newAlertItem, models.AlertValues{AlertType: alertType, PriceTrigger: priceTrigger}, emailRecipient)
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
		responseConfirmationString = fmt.Sprintf("The alert for item %s has been created successfully.\n The alert type is %d, the price trigger is %d and the email recipient is %s", newAlertItem, alertType, priceTrigger, emailRecipient)
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

func GetAllDefinedAlertsByItemHandler(ctx context.Context, alertsDefined store.AlertDefStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//Create key-value pairs (map) from the URL
		vars := mux.Vars(r)

		//Obtain from the url which item we are interested in reporting
		itemsToReport := vars["item"]

		dataAlertsDefined, err := alertsDefined.GetAlertsByItem(itemsToReport)
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

func CreateNewAlertRequestBodyParameters(r *http.Request) (alertType models.AlertType, priceTrigger int, emailRecipient string, err error) {

	//Variables to store the decoded string from message body
	var decodedString struct {
		alertType      string
		priceTrigger   string
		emailRecipient string
	}

	//Conversion handling variables
	var alertTypeInt int

	// Read the entire body
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return 0, 0, "", err
	}

	//Decode the parameters from the request body - url encoded string
	decodedValues, err := url.ParseQuery(string(bodyBytes))
	if err != nil {
		return 0, 0, "", err
	}

	//Extract the alertType and priceTrigger from the decoded values
	decodedString.alertType = decodedValues.Get("alertType")
	decodedString.priceTrigger = decodedValues.Get("priceTrigger")
	decodedString.emailRecipient = decodedValues.Get("emailRecipient")

	//Convert alertType from string to AlertType enum
	alertTypeInt, err = strconv.Atoi(decodedString.alertType)
	if err != nil {
		return 0, 0, "", fmt.Errorf("invalid alert type format")
	}
	alertType = models.AlertType(alertTypeInt)

	//Convert priceTrigger from string to int
	priceTrigger, err = strconv.Atoi(decodedString.priceTrigger)
	if err != nil {
		return 0, 0, "", fmt.Errorf("invalid new alert details format")
	}

	// Set emailRecipient
	emailRecipient = decodedString.emailRecipient

	// Return the extracted and converted values
	return alertType, priceTrigger, emailRecipient, nil
}

func CreateNewAlertObtainItemFromURL(r *http.Request) (item string, err error) {

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) >= 4 {
		item = parts[2]
	}

	return item, nil
}
