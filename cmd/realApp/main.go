package main

import (
	"Price_Notification_System/api"
	"Price_Notification_System/config"
	"Price_Notification_System/output"
	"Price_Notification_System/producer/trades"
	"Price_Notification_System/store"
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

func main() {
	var (
		mainCtx          context.Context
		cancel           context.CancelFunc
		eg               *errgroup.Group
		ctx              context.Context
		objects          []string
		individualTrades chan trades.TradeItems
		itemTradeHistory store.TradeStore
		alertStore       store.AlertDefStore
		errNewDBAlert    error
	)

	//create slice of objects that will be traded
	objects = []string{"Iron Man Figure", "Hulk Figure", "Deadpool Figure", "Wolverine Figure", "Spider-Man Figure",
		"Thor Figure", "Superman Figure", "Batman Figure", "Wonder-Woman Figure", "Captain America Figure"}

	//Create variables that will hold the individual trades as the TradeItems type from the trades package
	individualTrades = make(chan trades.TradeItems)

	//Load any required environment variables
	enVariables, errLoadingVariables := config.LoadEnvVariables()
	if errLoadingVariables != nil {
		log.Fatal("error loading environment variables: %v", errLoadingVariables)
	}
	//Load the DB connection string
	dbConnStr := config.LoadDBConnectionStr(enVariables)

	//Create instances of the HistoricalData/ Alerts store
	//itemTradeHistory = store.NewInMemoryTradeStore()
	//alertStore = store.NewInMemoryAlertStore()

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

	//Generate a 'trade' 'randomly' between 1-5 seconds
	eg.Go(func() error {
		return tradeTrigger(ctx, objects, individualTrades)
	})

	//Call the output function to process the trade
	eg.Go(func() error {
		return output.Outputs(ctx, individualTrades, itemTradeHistory, alertStore, os.Stdout)
	})

	//Run the HTTP server to allow API connections
	//ctxHTTPServer := context.Background()
	eg.Go(func() error { return api.HTTPServer(ctx, itemTradeHistory, alertStore) })

	//call method `Wait()` to ensure the program waits for all goroutines to complete
	err := eg.Wait()

	//output whether any errors occurred
	if err != nil {
		log.Fatal("Error:", err)
	} else {
		fmt.Println("All trades processed")
	}

}

// Function that triggers a set amount of trades (equal to i max value).
// trades are triggered 'randomly' between 1 and 5 second intervals.
func tradeTrigger(ctx context.Context, objects []string, individualTrades chan trades.TradeItems) error {

	defer close(individualTrades) // Close the channel to signal no more trades will be sent

	for i := 0; i < 300; i++ {
		randomSecs := int((rand.Float64() * 4.0) + 1)

		select {

		case <-ctx.Done():
			fmt.Println("Context cancelled, stopping trade trigger")

			return ctx.Err()

		default:
			time.Sleep(time.Duration(randomSecs) * time.Second)

			errFromTrades := trades.Trade(ctx, objects, individualTrades)
			if errFromTrades != nil {
				return errFromTrades
			}
			fmt.Printf("Trade %d\n", i)
		}
	}

	return nil
}
