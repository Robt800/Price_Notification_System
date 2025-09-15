package output

import (
	"Price_Notification_System/models"
	"Price_Notification_System/producer/trades"
	"Price_Notification_System/store"
	"context"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestOutput(t *testing.T) {

	type testDef struct {
		instanceName          string
		dataToAddToChannel    []trades.TradeItems
		producedData          chan trades.TradeItems
		tradeStore            store.TradeStore
		write                 strings.Builder
		expectedWrittenOutput []string
		expectedTradeStore    store.TradeStore
		wantErr               error
	}

	var tests = []testDef{
		{
			instanceName: "Test 1 - Output 3 trades to a empty trade store",
			dataToAddToChannel: []trades.TradeItems{
				{Object: "Shelby GT500",
					Timestamp: time.Date(2023, 10, 1, 13, 50, 0, 0, time.UTC),
					Price:     700},
				{Object: "Mustang",
					Timestamp: time.Date(2023, 10, 30, 14, 50, 0, 0, time.UTC),
					Price:     800},
				{Object: "Corvette",
					Timestamp: time.Date(2023, 10, 31, 15, 50, 0, 0, time.UTC),
					Price:     900},
			},
			producedData: make(chan trades.TradeItems, 0),
			tradeStore:   store.NewInMemoryTradeStore(),
			write:        strings.Builder{},
			expectedWrittenOutput: []string{
				"{Shelby GT500 2023-10-01 13:50:00 +0000 UTC 700}",
				"{Mustang 2023-10-30 14:50:00 +0000 UTC 800}",
				"{Corvette 2023-10-31 15:50:00 +0000 UTC 900}",
			},
			expectedTradeStore: store.NewInMemoryTradeStoreWithData(&map[time.Time]models.HistoricalDataValues{
				time.Date(2023, 10, 1, 13, 50, 0, 0, time.UTC):  {Object: "Shelby GT500", Price: 700},
				time.Date(2023, 10, 30, 14, 50, 0, 0, time.UTC): {Object: "Mustang", Price: 800},
				time.Date(2023, 10, 31, 15, 50, 0, 0, time.UTC): {Object: "Corvette", Price: 900},
			}),
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.instanceName, func(t *testing.T) {
			// Create a context
			ctx := context.Background()

			// Ensure channel is closed after use
			defer close(tt.producedData)

			// Add trade data to the channel
			go func() {
				for _, trade := range tt.dataToAddToChannel {
					tt.producedData <- trade
				}
			}()

			// Call the Outputs function
			err := Outputs(ctx, tt.producedData, tt.tradeStore, &tt.write, nil)

			// Split the strings.Builder into a slice of strings
			writeSliceStrings := splitStringsBuilder(tt.write)

			// Check for errors
			if err != nil && err != tt.wantErr {
				t.Errorf("Outputs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Check if the write buffer contains the expected output
			for i, writeSliceElement := range writeSliceStrings {
				if writeSliceElement != tt.expectedWrittenOutput[i] {
					t.Errorf("Outputs() got = %v, want %v", writeSliceElement, tt.expectedWrittenOutput[i])
				}
			}

			// Check if the trade store contains the expected trades
			if !reflect.DeepEqual(tt.tradeStore, tt.expectedTradeStore) {
				t.Errorf("TradeStore() got = %v, want %v", tt.tradeStore, tt.expectedTradeStore)
			}
		})
	}
}

func splitStringsBuilder(sb strings.Builder) []string {
	// Split the strings.Builder into a slice of strings
	sbSplit := strings.Split(sb.String(), "\n")
	return sbSplit[:len(sbSplit)-1] // Exclude the last empty string
}
