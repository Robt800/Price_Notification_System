package trades

import (
	"context"
	"encoding/json"
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
// Also the main function is located within the tradeImpl function - the Trade function acts as a wrapper to this to
// allow easier unit testing of the timestamp
func Trade(ctx context.Context, tradeObjects []string, individualTrades chan []byte) error {
	return tradeImpl(ctx, tradeObjects, individualTrades, func() time.Time { return time.Now() })
}

func tradeImpl(ctx context.Context, tradeObjects []string, individualTrades chan []byte, nowProvider func() time.Time) error {
	//Create a 'random' trade
	TradedItem := randomTrade(tradeObjects, nowProvider)

	//Marshall the traded struct into JSON format
	TradedItemJSON, err := json.Marshal(TradedItem)
	if err != nil {
		return err
	}

	//return TradedItemJSON
	individualTrades <- TradedItemJSON

	return nil
}

func randomTrade(tradeObject []string, nowProvider func() time.Time) (TradedItem TradeItems) {
	randObjectNo := rand.Intn(len(tradeObject))
	randPrice := rand.Intn(250) + 800 //generate random price between £8 & £10.50

	//Create a 'random' trade based on the above
	TradedItem = TradeItems{
		Object:    tradeObject[randObjectNo],
		Timestamp: nowProvider(),
		Price:     randPrice,
	}
	return TradedItem
}
