package api

import (
	store "Price_Notification_System/producer/store"
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HTTPServer(ctx context.Context, itemTradeHistory store.HistoricalData) error {

	//Create a mux router instance which can be used assign routes to etc.
	r := mux.NewRouter()

	//Define the routes
	r.HandleFunc("/item_traded/{id}", GetItemPriceHandler(itemTradeHistory)).Methods("GET")

	//Start the server
	log.Fatal(http.ListenAndServe(":8080", r))

	return nil
}
