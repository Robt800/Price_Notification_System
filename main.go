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

	//Create variables that will hold the individual trades and all trades together
	//as a slice of byte - which JSON format uses to store data
	var individualTrades, allTrades []byte

	//Create a channel of single bool - used to trigger trades
	TriggerChannel := make(chan bool)

	//Generate a 'true' onto the channel randomly between 1-5 seconds - i.e. send to the channel
	go tradeTrigger(TriggerChannel)

	//use range function to trigger call trade function.
	//When the trade has been completed - print the trade out and
	//add the trade to the 'allTrades' slice that holds 'all trades'
	for _ = range TriggerChannel {
		individualTrades = Trades.Trade(Objects)
		fmt.Printf("%v\n", string(individualTrades))
		allTrades = append(allTrades, individualTrades...)
	}

	fmt.Printf("%v\n", string(allTrades))

}

// Function that triggers a set amount of trades (equal to i max value).
// Trades are triggered 'randomly' between 1 and 5 second intervals.
func tradeTrigger(trigger chan<- bool) {
	for i := 0; i < 20; i++ {
		randomSecs := int((rand.Float64() * 4.0) + 1)
		time.Sleep(time.Duration(randomSecs) * time.Second)
		trigger <- true
	}
	close(trigger)
}
