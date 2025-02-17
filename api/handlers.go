package api

import (
	"Price_Notification_System/trades"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetItemPriceHandler(historicalTransactions *[]trades.TradeItems) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Create key-value pairs (map) from the URL
		vars := mux.Vars(r)

		//Obtain from the url which item we are interested in reporting
		itemsToReport := vars["id"]

		//Create a var to store the results in
		var results []trades.TradeItems

		//Range over the shared data store and store relevant transactions within the response slice
		for _, v := range *historicalTransactions {
			if v.Object == itemsToReport {
				results = append(results, v)
			}
		}

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
