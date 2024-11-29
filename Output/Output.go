package Output

import (
	"fmt"
	"sync"
	"time"
)

// Outputs ensures the data from the channel (i.e. the trade) is genuine - if it is, it prints it
func Outputs(producedData chan []byte, wg *sync.WaitGroup) {
	done := false
	for !done {
		select {
		case tradeData, ok := <-producedData:
			if !ok {
				done = true
			}
			fmt.Printf("%v\n", string(tradeData))
		case <-time.After(time.Second * 10):
			wg.Done()
		}
	}
	wg.Done()
}
