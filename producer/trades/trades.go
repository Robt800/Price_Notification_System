package trades

import (
	"context"
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
// Also the main function is located within the tradeImpl function - the Trade function acts as a wrapper to this to
// allow easier unit testing of the timestamp
func Trade(ctx context.Context, tradeObjects []string, individualTrades chan TradeItems) error {
	return tradeImpl(ctx, tradeObjects, individualTrades, func() time.Time { return time.Now() })
}

func tradeImpl(ctx context.Context, tradeObjects []string, individualTrades chan TradeItems, nowProvider func() time.Time) error {
	//Create a 'random' trade
	TradedItem := randomTrade(tradeObjects, nowProvider)

	//return TradedItem
	individualTrades <- TradedItem

	if ctx.Err() != nil {
		fmt.Printf("Context cancelled with error:%v\n", ctx.Err())
		return ctx.Err()
	}
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
