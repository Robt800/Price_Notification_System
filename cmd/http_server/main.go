package main

import (
	"Price_Notification_System/api"
	"Price_Notification_System/config"
	"Price_Notification_System/store"
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"os/signal"
)

func main() {
	var (
		mainCtx          context.Context
		cancel           context.CancelFunc
		eg               *errgroup.Group
		ctx              context.Context
		itemTradeHistory store.TradeStore
		alertStore       store.AlertDefStore
		errNewDBAlert    error
	)

	//Load any required environment variables
	enVariables, errLoadingVariables := config.LoadEnvVariables()
	if errLoadingVariables != nil {
		log.Fatal("error loading environment variables: %v", errLoadingVariables)
	}
	//Load the DB connection string
	dbConnStr := config.LoadDBConnectionStr(enVariables)

	itemTradeHistory, errNewDBTrade := store.NewDBTradeStore(dbConnStr)
	if errNewDBTrade != nil {
		log.Fatal("error creating trade store:", errNewDBTrade)
	}

	alertStore, errNewDBAlert = store.NewDBAlertStore(dbConnStr)
	if errNewDBAlert != nil {
		log.Fatal(errNewDBAlert)
	}

	//mainCtx instance to store the context which will be used - // Set up a context that cancels when you hit Ctrl+C:
	mainCtx, cancel = signal.NotifyContext(context.Background(), os.Interrupt)

	//cancel function is called as final part of program to release resources associated with the context when the function returns
	defer cancel()

	//errgroup and context variables created from the errgroup package.  Used to sync & error propagate between goroutines
	eg, ctx = errgroup.WithContext(mainCtx)

	//Run the HTTP server to allow API connections
	//ctxHTTPServer := context.Background()
	eg.Go(func() error { return api.HTTPServer(ctx, itemTradeHistory, alertStore) })

	//call method `Wait()` to ensure the program waits for all goroutines to complete
	err := eg.Wait()

	//output whether any errors occurred
	if err != nil {
		log.Fatal("Error: " + err.Error() + "\n")
	} else {
		fmt.Println("HTTP server stopped gracefully\n")
	}
}
