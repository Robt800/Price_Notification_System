package Trades

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type TradeItems struct {
	Object    string
	Timestamp time.Time
	Price     int
}

// Trade function that
func Trade(TradeObject, TradeObjectBuffer []string) []byte {
	//If the TradeObject slice is empty (because all previously moved to buffer) - empty buffer back to main slice
	if len(TradeObject) == 0 {
		TradeObject = append(TradeObject, TradeObjectBuffer...)
		TradeObjectBuffer = append(TradeObjectBuffer[:0])
	}

	randObjectNo := rand.Intn(len(TradeObject))
	randPrice := rand.Intn(250) + 800 //generate random price between £8 & £10.50

	//Create a 'random' trade based on the above
	TradedItem := TradeItems{
		Object:    TradeObject[randObjectNo],
		Timestamp: time.Now(),
		Price:     randPrice,
	}

	//Remove the traded item to a buffer so duplicates don't occur until all items have been traded at least once
	for i := 0; i < len(TradeObjectBuffer); i++ {
		if TradeObjectBuffer[i] == "" {
			TradeObjectBuffer[i] = TradeObject[randObjectNo]
			break
		}
	}

	//TradeObjectBuffer = append(TradeObjectBuffer, TradeObject[randObjectNo])
	TradeObject = append(TradeObject[:randObjectNo], TradeObject[randObjectNo+1:]...) //UPTO!!!!!!!!! - need to fix this

	//Marshall the traded struct into JSON format
	TradedItemJSON, err := json.Marshal(TradedItem)
	if err != nil {
		fmt.Println(err)
	}

	return TradedItemJSON
}
