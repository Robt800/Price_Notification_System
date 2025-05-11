package api

import (
	store "Price_Notification_System/store"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func HTTPServer(ctx context.Context, itemTradeHistory store.TradeStore) error {

	//Create a mux router instance which can be used assign routes to etc.
	r := mux.NewRouter()

	//Define the routes
	r.HandleFunc("/item_traded/{id}", GetTradesByItemHandler(ctx, itemTradeHistory)).Methods("GET")

	//Start the server
	//log.Fatal(http.ListenAndServe(":8080", r))

	if ctx.Err() != nil {
		fmt.Printf("Context error:%v", ctx.Err())
		return ctx.Err()
	}

	httpError := http.ListenAndServe(":8080", r)

	if httpError != nil {
		fmt.Printf("HTTP server error:%v", httpError)
		return httpError
	}

	return nil
	//return nil
}
