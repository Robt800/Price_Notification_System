package api

import (
	"context"
	"github.com/gorilla/mux\"
	"log"
	"net/http"
)

func HTTPServer(ctx context.Context) error {

	//Create a mux router instance which can be used assign routes to etc.
	r := mux.NewRouter()

	//Define the routes
	r.HandleFunc("/item_traded/{id}", GetItemPriceHandler).Methods("GET")

	//Start the server
	log.Fatal(http.ListenAndServe(":8080", r))

	return nil
}
