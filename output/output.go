package output

import (
	"Price_Notification_System/trades"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type alert struct {
	alertType    string //"Price Alert - Low Price", "Price Alert - High Price"
	object       string
	priceTrigger int
}

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

func OutputsWithNotification(ctx context.Context, producedData chan []byte) error {

	//Obtain the produced data from the channel & call 'processTradeFromChannel'
	var actualTrade []byte
	var ok bool

	done := false
	for !done {
		select {
		case actualTrade, ok = <-producedData:
			if !ok {
				done = true
			}
			err := processTradeFromChannel(ctx, actualTrade)
			if err != nil {
				return err
			}
		case <-time.After(time.Second * 10):
			done = true
		}
	}
	return nil
}

func processTradeFromChannel(ctx context.Context, actualTrade []byte) error {
	var (
		alert1                    alert
		alertNeeded               bool
		alertGenerated            chan []byte
		alertFromChannel          []byte
		alertFromChannelUnmarshal trades.TradeItems
	)

	//Unmarshall the trade into the principle elements for easier comparison
	var tradedItem trades.TradeItems

	err := json.Unmarshal(actualTrade, &tradedItem)
	if err != nil {
		return err
	}

	//TEMPORARILY output minor details of the trade - used for testing - delete once happy - #TODO - delete once tested
	fmt.Printf("The trade of %v was made at a price of %v\n", tradedItem.Object, tradedItem.Price)

	//Create an instance of an alert
	alert1 = alert{
		alertType:    "Price Alert - High Price",
		object:       "Hulk Figure",
		priceTrigger: 865,
	}

	//Determine if an alert is required
	alertNeeded = alertRequired(ctx, tradedItem, alert1)

	//Create the channel to store the data
	alertGenerated = make(chan []byte, 1)
	defer close(alertGenerated)

	//Obtain details of the alert and place in a channel
	if alertNeeded {
		alertGenerated <- actualTrade
	}

	//Consume the data from the alertGenerated channel
	switch {
	case alertNeeded:
		alertFromChannel = <-alertGenerated

		err = json.Unmarshal(alertFromChannel, &alertFromChannelUnmarshal)
		if err != nil {
			return err
		}
		fmt.Printf("The following alert has been generated:\n Alert type: %v\n Details of the trade matching this alert: %v\n", alert1.alertType, alertFromChannelUnmarshal)

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
