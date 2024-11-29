package main

import (
	"Price_Notification_System/Output"
	"Price_Notification_System/Trades"
	"math/rand"
	"sync"
	"time"
)

func main() {
	//create slice of objects that will be traded
	Objects := []string{"Iron Man Figure", "Hulk Figure", "Deadpool Figure", "Wolverine Figure", "Spider-Man Figure",
		"Thor Figure", "Superman Figure", "Batman Figure", "Wonder-Woman Figure", "Captain America Figure"}

	//Create variables that will hold the individual trades as a slice of byte - which JSON format uses to store data
	//var wg used as a waitgroup to ensure no deadlock/ run conditions on relevant channels
	individualTrades := make(chan []byte)
	var wg sync.WaitGroup

	//Create a channel of single bool - used to trigger trades
	TriggerChannel := make(chan bool)

	//Generate a 'true' onto the channel randomly between 1-5 seconds - i.e. send to the channel
	go tradeTrigger(TriggerChannel)

	//use range function to trigger call trade function.
	//When the trade has been completed - print the trade out and
	for _ = range TriggerChannel {
		wg.Add(1)
		go Trades.Trade(Objects, individualTrades, &wg)
	}

	//Call the Output function to process the trade
	wg.Add(1)
	go Output.Outputs(individualTrades, &wg)

	//Use wg.Wait function to wait until previous go routines have completed before finishing the program
	wg.Wait()
	close(individualTrades)

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
