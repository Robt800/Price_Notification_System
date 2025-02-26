package output

import (
	store "Price_Notification_System/producer/store"
	"Price_Notification_System/producer/trades"
	"context"
	"fmt"
	"time"
)

// Outputs ensures the data from the channel (i.e. the trade) is genuine - if it is, it prints it
func Outputs(ctx context.Context, producedData chan []byte) error {
	done := false
	for !done &&
		(ctx.Err() == nil) {
		select {
		case tradeData, ok := <-producedData:
			if !ok {
				done = true
			}
			fmt.Printf("%v\n", string(tradeData))
		case <-time.After(time.Second * 10):
			done = true
		}
	}
	return nil
}

func OutputsWithNotification(ctx context.Context, producedData chan trades.TradeItems, tradeHistory store.HistoricalData) error {

	//Obtain the produced data from the channel & call 'processTradeFromChannel'
	var actualTrade trades.TradeItems
	var ok bool

	done := false
	for !done {
		select {
		case actualTrade, ok = <-producedData:
			if !ok {
				done = true
			}
			err := processTradeFromChannel(ctx, actualTrade, tradeHistory)
			if err != nil {
				return err
			}
		case <-time.After(time.Second * 10):
			done = true
		}
	}
	return nil
}

func processTradeFromChannel(ctx context.Context, actualTrade trades.TradeItems, tradeHistory store.HistoricalData) error {
	var (
		alert1           alert
		alertNeeded      bool
		alertGenerated   chan trades.TradeItems
		alertFromChannel trades.TradeItems
	)

	//Add the trade to the history
	addTradeToHistory(tradeHistory, actualTrade)

	//TEMPORARILY output minor details of the trade - used for testing - delete once happy - #TODO - delete once tested
	//fmt.Printf("The trade of %v was made at a price of %v\n", actualTrade.Object, actualTrade.Price)
	fmt.Printf("The full trade history is %v items long\n", len(tradeHistory))

	//Determine if an alert is required
	alertNeeded = alertRequired(ctx, actualTrade, alert1)

	//Create the channel to store the data
	alertGenerated = make(chan trades.TradeItems, 1)
	defer close(alertGenerated)

	//Obtain details of the alert and place in a channel
	if alertNeeded {
		alertGenerated <- actualTrade
	}

	//Consume the data from the alertGenerated channel
	switch {
	case alertNeeded:
		alertFromChannel = <-alertGenerated

		fmt.Printf("The following alert has been generated:\n Alert type: %v\n Details of the trade matching this alert: %v\n", alert1.alertType, alertFromChannel)

	}

	return nil
}

func alertRequired(ctx context.Context, actualTrade trades.TradeItems, alertParams alert) bool {

	switch alertParams.alertType {
	case "Price Alert - Low Price":
		if (alertParams.object == actualTrade.Object) &&
			(actualTrade.Price <= alertParams.priceTrigger) {
			return true
		}
	case "Price Alert - High Price":
		if (alertParams.object == actualTrade.Object) &&
			(actualTrade.Price >= alertParams.priceTrigger) {
			return true
		}
	}
	return false
}

func addTradeToHistory(tradeHistory store.HistoricalData, trade trades.TradeItems) {
	tradeDataValues := store.HistoricalDataValues{Object: trade.Object, Price: trade.Price}

	tradeHistory.Add(trade.Timestamp, tradeDataValues)
}
