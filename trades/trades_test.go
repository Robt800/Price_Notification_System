package trades

import (
	"context"
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"math/rand"
	"slices"
	"testing"
	"time"
)

func TestTrade(t *testing.T) {
	// set random seed for reproducibility
	// this ensures that on every run, rand.Intn will produce the same sequence of numbers
	rand.New(rand.NewSource(1))

	type testDef struct {
		testInstanceName string
		tradeObjects     []string
		timeProvider     func() time.Time
		expected         TradeItems
		wantErr          bool
	}

	var tests = []testDef{
		testDef{testInstanceName: "successful trade1", tradeObjects: []string{"Laptop", "PC", "Monitor", "Keyboard", "Mouse"},
			timeProvider: func() time.Time { return time.Date(2024, time.January, 30, 16, 30, 0, 0, time.UTC) },
			expected:     TradeItems{Object: "Monitor", Timestamp: time.Date(2024, time.January, 30, 16, 30, 0, 0, time.UTC), Price: 950},
			wantErr:      false},
		testDef{testInstanceName: "successful trade2", tradeObjects: []string{"Table", "Chair", "Plate", "Knife", "Fork"},
			timeProvider: func() time.Time { return time.Date(2025, time.February, 25, 8, 15, 59, 0, time.UTC) },
			expected:     TradeItems{Object: "Chair", Timestamp: time.Date(2025, time.February, 25, 8, 15, 59, 0, time.UTC), Price: 980},
			wantErr:      false},
		testDef{testInstanceName: "successful trade3", tradeObjects: []string{"Corvette", "Ford", "Pontiac", "Dodge", "Cadillac"},
			timeProvider: func() time.Time { return time.Date(2023, time.November, 05, 19, 29, 3, 0, time.UTC) },
			expected:     TradeItems{Object: "Pontiac", Timestamp: time.Date(2023, time.November, 05, 19, 29, 3, 0, time.UTC), Price: 890},
			wantErr:      false},
	}

	// Iterate through the tests
	for _, tt := range tests {
		t.Run(tt.testInstanceName, func(t *testing.T) {

			ctx := context.Background()
			resultChannel := make(chan []byte, 1)
			defer close(resultChannel)

			err := tradeImpl(ctx, tt.tradeObjects, resultChannel, tt.timeProvider)
			if (err != nil) != tt.wantErr {
				t.Errorf("tradeImpl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				select {
				case trade := <-resultChannel:
					var tradedItem TradeItems

					if err := json.Unmarshal(trade, &tradedItem); err != nil {
						t.Errorf("Failed to unmarshal trade: %v", err)
					}
					if validObject := slices.Contains(tt.tradeObjects, tradedItem.Object); !validObject {
						t.Errorf("The object %v is not included within the allowed trade items %v", tradedItem.Object, tt.tradeObjects)
					}
					if timeDiff := cmp.Diff(tt.expected.Timestamp, tradedItem.Timestamp); timeDiff != "" {
						t.Errorf("The timestamp of the trade: %v does not match the anticipated timestamp: %v", tradedItem.Timestamp, tt.expected.Timestamp)
					}
					if priceWithinTolerance := (tradedItem.Price >= 800) && (tradedItem.Price <= 1050); !priceWithinTolerance {
						t.Errorf("The traded price of %v is outside the permissable range of 800 - 1050", tradedItem.Price)
					}
				}
			}
		})
	}
}
