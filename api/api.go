package api

import (
	"context"
	"github.com/gorilla/mux\"
)

func HTTPServer(ctx context.Context) error {

	//Create a mux router instance which can be used assign routes to etc.
	r := mux.NewRouter()

	//Define the routes
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/users", GetAllUsers).Methods("GET")

}
