package Output

import (
	"fmt"
	"sync"
	"time"
)

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
