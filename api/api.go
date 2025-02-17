package api

import (
	"Price_Notification_System/trades"
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HTTPServer(ctx context.Context, historicalTransactions *[]trades.TradeItems) error {

	//Create a mux router instance which can be used assign routes to etc.
	r := mux.NewRouter()

	//Define the routes
	r.HandleFunc("/item_traded/{id}", GetItemPriceHandler(historicalTransactions)).Methods("GET")

	//Start the server
	log.Fatal(http.ListenAndServe(":8080", r))

	return nil
}
