package api

import (
	store "Price_Notification_System/store"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func HTTPServer(ctx context.Context, itemTradeHistory store.TradeStore, alertsDefined store.AlertDefStore) error {

	//Create a mux router instance which can be used assign routes to etc.
	r := mux.NewRouter()

	//Define the routes
	r.HandleFunc("/items/trade-history/{item}", GetTradesByItemHandler(ctx, itemTradeHistory)).Methods("GET")

	r.HandleFunc("/items/{item}/alerts", CreateNewAlertHandler(ctx, alertsDefined)).Methods("POST") // Need to pass alertType & priceTrigger in request body, not URI

	r.HandleFunc("/items/all/alerts", GetAllDefinedAlertsHandler(ctx, alertsDefined)).Methods("GET")

	r.HandleFunc("/items/{item}/alerts", GetAllDefinedAlertsByItemHandler(ctx, alertsDefined)).Methods("GET") // Get all alerts for a specific item

	//Check if the context has been cancelled
	if ctx.Err() != nil {
		fmt.Printf("Context error:%v", ctx.Err())
		return ctx.Err()
	}

	//Start the server
	httpError := http.ListenAndServe(":8080", r)

	if httpError != nil {
		fmt.Printf("HTTP server error:%v", httpError)
		return httpError
	}

	return nil
	//return nil
}
