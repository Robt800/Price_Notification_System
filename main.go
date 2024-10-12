package main

import (
	"Price_Notification_System/Trades"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//create slice of objects that will be traded
	Objects := []string{"Iron Man Figure", "Hulk Figure", "Deadpool Figure", "Wolverine Figure", "Spider-Man Figure",
		"Thor Figure", "Superman Figure", "Batman Figure", "Wonder-Woman Figure", "Captain America Figure"}

	//Create buffer for objects that have been traded
	ObjectsBuffer := make([]string, len(Objects))

	var individualTrades, allTrades []byte //Create variables that will hold the individual trades and all trades together
	//as a slice of byte - which JSON format uses to store data

	//Create a channel of single bool - used to trigger trades
	TriggerChannel := make(chan bool)

	//Generate a 'true' onto the channel randomly between 1-5 seconds - i.e. send to the channel
	go tradeTrigger(TriggerChannel)

	//use range function to trigger call trade function
	for _ = range TriggerChannel {
		individualTrades = Trades.Trade(Objects, ObjectsBuffer)
		allTrades = append(allTrades, individualTrades...)
	}

	fmt.Printf("%v\n", string(allTrades))

}

func tradeTrigger(trigger chan<- bool) {
	for i := 0; i < 11; i++ {
		randomSecs := int((rand.Float64() * 4.0) + 1)
		time.Sleep(time.Duration(randomSecs) * time.Second)
		trigger <- true
	}
	close(trigger)
}
