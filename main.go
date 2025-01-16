package main

import (
	"Price_Notification_System/Output"
	"Price_Notification_System/Trades"
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"math/rand"
	"time"
)

var (
	mainCtx          context.Context
	cancel           context.CancelFunc
	eg               *errgroup.Group
	ctx              context.Context
	objects          []string
	individualTrades chan []byte
)

func main() {
	//create slice of objects that will be traded
	objects = []string{"Iron Man Figure", "Hulk Figure", "Deadpool Figure", "Wolverine Figure", "Spider-Man Figure",
		"Thor Figure", "Superman Figure", "Batman Figure", "Wonder-Woman Figure", "Captain America Figure"}

	//Create variables that will hold the individual trades as a slice of byte - which JSON format uses to store data
	individualTrades = make(chan []byte)

	//mainCtx instance to store the context which will be used - time of 100secs is allowed before context cancellation
	mainCtx, cancel = context.WithTimeout(context.Background(), 100000*time.Millisecond)

	//cancel function is called as final part of program to release resources associated with the context when the function returns
	defer cancel()

	//errgroup and context variables created from the errgroup package.  Used to sync & error propagate between goroutines
	eg, ctx = errgroup.WithContext(mainCtx)

	//Generate a 'trade' 'randomly' between 1-5 seconds
	eg.Go(func() error {
		return tradeTrigger()
	})

	//Call the Output function to process the trade
	eg.Go(func() error {
		return output.Outputs(individualTrades, ctx)
	})

	//call method `Wait()` to ensure the program waits for all goroutines to complete
	err := eg.Wait()

	//Output whether any errors occurred
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("All trades processed")
	}

	//close the channel
	close(individualTrades)

}

// Function that triggers a set amount of trades (equal to i max value).
// Trades are triggered 'randomly' between 1 and 5 second intervals.
func tradeTrigger() error {
	for i := 0; i < 30; i++ {
		randomSecs := int((rand.Float64() * 4.0) + 1)
		time.Sleep(time.Duration(randomSecs) * time.Second)

		errFromTrades := trades.Trade(objects, individualTrades, ctx)
		if errFromTrades != nil {
			return errFromTrades
		}
	}

	return nil
}
