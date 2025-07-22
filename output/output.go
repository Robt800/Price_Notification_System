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
func Outputs(ctx context.Context, producedData chan trades.TradeItems, tradeStore store.TradeStore, alertStore store.AlertDefStore, write io.Writer) error {
	done := false
	for !done &&
		(ctx.Err() == nil) {
		select {
		case tradeData, ok := <-producedData:
			if !ok {
				done = true
				return nil
			}
			_, _ = fmt.Fprintf(write, "%v\n", tradeData)

			// Check if the trade meets any alert criteria
			alerts, err := alertStore.GetAlertsByItem(tradeData.Object)
			if err != nil {
				fmt.Fprintf(write, "Error retrieving alerts for item %s: %v\n", tradeData.Object, err)
				return err
			}

			// Determine if the trade matches any alert criteria
			tradeMatchesAlertCriteria := tradeMeetsAlertCriteria(tradeData, alerts)

			if tradeMatchesAlertCriteria {
				errAddingTrade := tradeStore.AddTrade(tradeData.Timestamp, models.HistoricalDataValues{Object: tradeData.Object, Price: tradeData.Price})
				if errAddingTrade != nil {
					fmt.Fprintf(write, "Error storing trade data: %v\n", errAddingTrade)
				} else {
					fmt.Fprintf(write, "Trade data stored successfully!\n")
				}
			}
		case <-time.After(time.Second * 20):
			fmt.Fprintf(write, "No trades in the last 20 seconds\n")
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

func tradeMeetsAlertCriteria(tradeData trades.TradeItems, alerts []models.AlertsByItemReturned) bool {

	for _, alert := range alerts {

		if tradeData.Object == alert.Item &&
			alert.AlertType == 0 &&
			tradeData.Price <= alert.PriceTrigger {
			return true
		} else if tradeData.Object == alert.Item &&
			alert.AlertType == 1 &&
			tradeData.Price >= alert.PriceTrigger {
			return true
		}
	}
	return false
}
