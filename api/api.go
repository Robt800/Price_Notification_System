package api

import (
	store "Price_Notification_System/store"
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
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
	//if ctx.Err() != nil {
	//	fmt.Printf("Context error:%v", ctx.Err())
	//	return ctx.Err()
	//}

	//Start the server
	//httpError := http.ListenAndServe(":8080", r)

	//Create the server instance
	srv := &http.Server{Addr: ":8080", Handler: r}

	// Goroutine to handle context cancellation
	go func() {
		<-ctx.Done()
		fmt.Println("HTTP server received context cancellation; shutting down HTTP server")
		// give server up to 5 seconds to shut down
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if errShutdown := srv.Shutdown(shutdownCtx); errShutdown != nil {
			fmt.Printf("HTTP server Shutdown error: %v\n", errShutdown)
		}
	}()

	//listen and action server requests
	err := srv.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("HTTP server closed")
		return nil // expected after shutdown
	}
	return err // unexpected error
}

//// Wait for the server to close
//errHTTPServer := srv.ListenAndServe()
//if err == http.ErrServerClosed {
//	fmt.Println("HTTP server closed")
//	return nil // expected after shutdown
//}
//return errHTTPServer // unexpected error

//
//
//
//
//if httpError != nil {
//	fmt.Printf("HTTP server error:%v", httpError)
//	return httpError
//}

//return nil
//}
