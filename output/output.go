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

	//Obtain the produced data from the channel
	var actualTrade []byte
	var ok bool

	select {
	case actualTrade, ok = <-producedData:
		if !ok {
			return nil
		}
	case <-time.After(time.Second * 10):
		return nil
	}

	//Unmarshall the trade into the principle elements for easier comparison
	var tradedItem trades.TradeItems

	err := json.Unmarshal(actualTrade, &tradedItem)
	if err != nil {
		return err
	}

	//TEMPORARILY output minor details of the trade
	fmt.Printf("The trade of %v was made at a price of %v", tradedItem.Object, tradedItem.Price)

	//Create an instance of an alert
	alert1 := alert{
		alertType:    "Price Alert - High Price",
		object:       "Hulk Figure",
		priceTrigger: 865,
	}

	//Determine if an alert is required
	alertNeeded := alertRequired(ctx, tradedItem, alert1)

	//Obtain details of the alert - TODO!!!!
	if alertNeeded {
		fmt.Println(string(actualTrade))
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
