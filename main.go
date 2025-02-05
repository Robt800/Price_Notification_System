package main

import (
	"Price_Notification_System/output"
	"Price_Notification_System/trades"
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"math/rand"
	"time"
)

func main() {
	var (
		mainCtx          context.Context
		cancel           context.CancelFunc
		eg               *errgroup.Group
		ctx              context.Context
		objects          []string
		individualTrades chan []byte
	)

	//create slice of objects that will be traded
	objects = []string{"Iron Man Figure", "Hulk Figure", "Deadpool Figure", "Wolverine Figure", "Spider-Man Figure",
		"Thor Figure", "Superman Figure", "Batman Figure", "Wonder-Woman Figure", "Captain America Figure"}

	//Create variables that will hold the individual trades as a slice of byte - which JSON format uses to store data
	individualTrades = make(chan []byte)
	defer close(individualTrades)

	//mainCtx instance to store the context which will be used - time of 100secs is allowed before context cancellation
	mainCtx, cancel = context.WithTimeout(context.Background(), 100000*time.Millisecond)

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
		return output.Outputs(ctx, individualTrades)
	})

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
func tradeTrigger(ctx context.Context, objects []string, individualTrades chan []byte) error {
	for i := 0; i < 30; i++ {
		randomSecs := int((rand.Float64() * 4.0) + 1)
		time.Sleep(time.Duration(randomSecs) * time.Second)

		errFromTrades := trades.Trade(ctx, objects, individualTrades)
		if errFromTrades != nil {
			return errFromTrades
		}
	}
	return nil
}
