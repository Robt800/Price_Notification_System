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
func Trade(tradeObjects, tradeObjectsBuffer []string) (tradeObjectsProcessed, tradeObjectsBufferProcessed []string,
	TradedItemJSON []byte) {

	//If the TradeObject slice is empty (because all previously moved to buffer) - empty buffer back to main slice
	tradeObjects, tradeObjectsBuffer = emptyBufferTradesToMain(tradeObjects, tradeObjectsBuffer)

	//Create a 'random' trade
	TradedItem, randomObjNumber := randomTrade(tradeObjects)

	//Copy traded item to buffer
	tradeObjects, tradeObjectsBuffer = moveTradeObjectToBuffer(tradeObjects, tradeObjectsBuffer, randomObjNumber)

	//Remove the traded item from the 'TradeObjects' so duplicates don't occur until all items have been traded at least once
	tradeObjects = removeTradedObject(tradeObjects, randomObjNumber)

	//Marshall the traded struct into JSON format
	TradedItemJSON, err := json.Marshal(TradedItem)
	if err != nil {
		fmt.Println(err)
	}

	tradeObjectsProcessed = tradeObjects
	tradeObjectsBufferProcessed = tradeObjectsBuffer

	return tradeObjectsProcessed, tradeObjectsBufferProcessed, TradedItemJSON
}

func emptyBufferTradesToMain(tradeObjects, tradeObjectsBuffer []string) (tradeObjectsProcessed, tradeObjectsBufferProcessed []string) {
	//If the TradeObject slice is empty (because all previously moved to buffer) - empty buffer back to main slice
	if len(tradeObjects) == 0 {
		tradeObjectsProcessed = append(tradeObjects, tradeObjectsBuffer...)
		tradeObjectsBufferProcessed = append(tradeObjectsBuffer[:0])
	} else {
		tradeObjectsProcessed = tradeObjects
		tradeObjectsBufferProcessed = tradeObjectsBuffer
	}
	return tradeObjectsProcessed, tradeObjectsBufferProcessed
}

func randomTrade(tradeObject []string) (TradedItem TradeItems, randObjectNo int) {
	randObjectNo = rand.Intn(len(tradeObject))
	randPrice := rand.Intn(250) + 800 //generate random price between £8 & £10.50

	//Create a 'random' trade based on the above
	TradedItem = TradeItems{
		Object:    tradeObject[randObjectNo],
		Timestamp: time.Now(),
		Price:     randPrice,
	}
	return TradedItem, randObjectNo
}

func moveTradeObjectToBuffer(tradeObjects, tradeObjectsBuffer []string, randObjectNo int) (tradeObjectsProcessed, tradeObjectsBufferProcessed []string) {
	for i := 0; i < len(tradeObjectsBuffer); i++ {
		if tradeObjectsBuffer[i] == "" {
			tradeObjectsBuffer[i] = tradeObjects[randObjectNo]
			break
		}
	}
	tradeObjectsProcessed = tradeObjects
	tradeObjectsBufferProcessed = tradeObjectsBuffer
	return tradeObjectsProcessed, tradeObjectsBufferProcessed

}

func removeTradedObject(tradeObjects []string, randObjectNo int) []string {
	return append(tradeObjects[:randObjectNo], tradeObjects[randObjectNo+1:]...)
}
