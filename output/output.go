package output

import (
	"context"
	"fmt"
	"time"
)

// Outputs ensures the data from the channel (i.e. the trade) is genuine - if it is, it prints it
func Outputs(ctx context.Context, producedData chan []byte) error {
	done := false
	for !done &&
		(ctx.Err() == nil) {
		select {
		case tradeData, ok := <-producedData:
			if !ok {
				done = true
			}
			fmt.Printf("%v\n", string(tradeData))
		case <-time.After(time.Second * 10):
			done = true
		}
	}
	return nil
}
