package output

import (
	"Price_Notification_System/models"
	"Price_Notification_System/producer/trades"
	store "Price_Notification_System/store"
	"context"
	"fmt"
	"io"
	"time"
)

// Outputs ensures the data from the channel (i.e. the trade) is genuine - if it is, it prints it
func Outputs(ctx context.Context, producedData chan trades.TradeItems, tradeStore store.TradeStore, write io.Writer) error {
	done := false
	for !done &&
		(ctx.Err() == nil) {
		select {
		case tradeData, ok := <-producedData:
			if !ok {
				done = true
			}
			fmt.Fprintf(write, "%v\n", tradeData)
			//fmt.Printf("%v\n", tradeData)
			tradeStore.AddTrade(tradeData.Timestamp, models.HistoricalDataValues{Object: tradeData.Object, Price: tradeData.Price})
		case <-time.After(time.Second * 10):
			fmt.Fprintf(write, "No trades in the last 10 seconds\n")
			done = true
		case <-ctx.Done():
			fmt.Fprintf(write, "Context cancelled\n")
			fmt.Fprintf(write, "The error was: %v\n", ctx.Err())
			done = true
			return ctx.Err()
		}
	}
	return nil
}
