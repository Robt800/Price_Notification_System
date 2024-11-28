package Trades

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

// TradeItems type that will store the 3 elements of data associated with each trade
type TradeItems struct {
	Object    string
	Timestamp time.Time
	Price     int
}

// Trade function that creates a random trade from the 'tradeObjects' that have been passed to it,
// marshalls this data into JSON format and return this from the function
func Trade(tradeObjects []string) (TradedItemJSON []byte) {

	//Create a 'random' trade
	TradedItem := randomTrade(tradeObjects)

	//Marshall the traded struct into JSON format
	TradedItemJSON, err := json.Marshal(TradedItem)
	if err != nil {
		fmt.Println(err)
	}

	return TradedItemJSON
}

func randomTrade(tradeObject []string) (TradedItem TradeItems) {
	randObjectNo := rand.Intn(len(tradeObject))
	randPrice := rand.Intn(250) + 800 //generate random price between £8 & £10.50

	//Create a 'random' trade based on the above
	TradedItem = TradeItems{
		Object:    tradeObject[randObjectNo],
		Timestamp: time.Now(),
		Price:     randPrice,
	}
	return TradedItem
}
