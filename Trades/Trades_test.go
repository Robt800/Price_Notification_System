package trades

import (
	"context"
	"errors"
	"fmt"
	"math"
	"slices"
	"sync"
	"testing"
	"time"
)

type args struct {
	tradeObjects     []string
	individualTrades chan []byte
	ctx              context.Context
}

var tests = []args{
	args{tradeObjects: []string{"Laptop", "PC", "Monitor", "Keyboard", "Mouse"}, individualTrades: make(chan []byte, 5), ctx: context.Background()},
	args{tradeObjects: []string{"Table", "Chair", "Plate", "Knife", "Fork"}, individualTrades: make(chan []byte, 5), ctx: context.Background()},
	args{tradeObjects: []string{"Corvette", "Ford", "Pontiac", "Dodge", "Cadillac"}, individualTrades: make(chan []byte, 5), ctx: context.Background()},
}

type returnedData struct {
	fullTrade  string
	item       string
	dateStamp  string
	price      string
	priceAsInt int
}

var returnedDataSets = make([]returnedData, 3)

func TestTrade(t *testing.T) {
	var wg sync.WaitGroup

	for i := range tests {
		wg.Add(1)
		go func(tt *args) {
			defer wg.Done()
			defer close(tt.individualTrades)
			err := Trade(tt.tradeObjects, tt.individualTrades, tt.ctx)
			if err != nil {
				t.Error(err)
			}
		}(&tests[i])
	}

	wg.Wait()

	//Get elements from the trade
	fmt.Println("The returned data from the function is as follows:")
	for i, test := range tests {
		for trade := range test.individualTrades {
			fmt.Printf("%s\n", trade)
			returnedDataSets[i].item = obtainObjectTraded(string(trade))
			returnedDataSets[i].dateStamp = obtainTimestampTrade(string(trade))
			returnedDataSets[i].price = obtainPriceTrade(string(trade))
		}
	}

	//Test to see if the object traded is valid
	for i, data := range returnedDataSets {
		if !slices.Contains(tests[i].tradeObjects, data.item) {
			t.Error("Item not found in tradedObjects, expected something from:", tests[i].tradeObjects,
				"got:", data.item)
		}
	}

	//Test to see if the timestamp looks reasonable
	timeStampOK, timeStampOfTrade, err := testTimeStampOK()
	if err != nil {
		t.Error(err)
	}
	if !timeStampOK {
		t.Error("The timestamp of the trade is not within tolerance. i.e. ", timeStampOfTrade)
	}

	//Test to see if the price looks reasonable
	for i := 0; i < len(returnedDataSets); i++ {
		returnedDataSets[i].priceAsInt, err = strToInt(returnedDataSets[i].price)
		if err != nil {
			t.Error(err)
		}
		if returnedDataSets[i].priceAsInt < 800 ||
			returnedDataSets[i].priceAsInt > 1050 {
			t.Error("The returned price is outside the expected of £8 to £10.50")
		}
	}

}

func obtainObjectTraded(fullTrade string) string {
	startPos := 11
	endPos := elementTradedDoubleQuotesPos(fullTrade, 4)
	return fullTrade[startPos:endPos]
}

func elementTradedDoubleQuotesPos(fullTrade string, doubleQuoteCountPosReq int) int {
	noDoubleQuotesFound := 0
	i := 0
	var c rune = 0
	for i, c = range fullTrade {
		if c == rune('"') {
			noDoubleQuotesFound++
		}
		if noDoubleQuotesFound == doubleQuoteCountPosReq {
			break
		}
	}
	return i
}

func obtainTimestampTrade(fullTrade string) string {
	startPos := elementTradedDoubleQuotesPos(fullTrade, 7) + 1
	endPos := elementTradedDoubleQuotesPos(fullTrade, 8)
	return fullTrade[startPos:endPos]
}

func obtainPriceTrade(fullTrade string) string {
	startPos := elementTradedDoubleQuotesPos(fullTrade, 10) + 2
	endPos := len(fullTrade) - 1
	return fullTrade[startPos:endPos]
}

func getTimeStampFromString(date string) (time.Time, error) {
	const layout = "2006-01-02T15:04:05"

	dateTime, err := time.Parse(layout, date[:len(date)-1])
	return dateTime, err
}

func testTimeStampOK() (bool, time.Time, error) {
	var timeStampOK bool
	var timeStampOfTrade time.Time
	for _, data := range returnedDataSets {
		timestampOfTrade, err := getTimeStampFromString(data.dateStamp)
		if err != nil {
			return false, timestampOfTrade, err
		}

		timeDiff := time.Now().Sub(timestampOfTrade)
		timeDiffSeconds := timeDiff.Seconds()
		timeStampOK = timeDiffSeconds < 5
	}
	return timeStampOK, timeStampOfTrade, nil
}

func strToInt(priceAsString string) (int, error) {
	intsMap := map[string]int{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}

	var strValAsInt int

	for i, v := range priceAsString {
		val, exists := intsMap[string(v)]
		if !exists {
			return 0, errors.New("The price is not a numerical value")
		}
		posMultiplier := int(math.Pow10(len(priceAsString) - i - 1))
		strValAsInt = (posMultiplier * val) + strValAsInt
	}
	return strValAsInt, nil
}
