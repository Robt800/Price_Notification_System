package api

import (
	"Price_Notification_System/models"
	"Price_Notification_System/service"
	store "Price_Notification_System/store"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
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
			log.Fatal("The HTTP server failed to get the results due to error: ", err)
		}

		////Range over the shared data store and store relevant transactions within the response slice
		//for k, v := range itemTradeHistory {
		//	if v.Object == itemsToReport {
		//		results[k] = v
		//	}
		//}

		//Convert the results to JSON in readiness to respond with the results
		resultsJSON, err := json.Marshal(results)
		if err != nil {
			log.Fatal("The HTTP server failed to get the results due to error: ", err)
		}

		//Write the results to the
		_, errWrite := w.Write(resultsJSON)
		if errWrite != nil {
			log.Fatal("The HTTP server failed to return the results due to error: ", errWrite)
		}
	}
}
