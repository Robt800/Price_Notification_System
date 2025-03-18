package models

import (
	"errors"
	"time"
)

// HistoricalDataValues type definition - used to define the elements of each trade
type HistoricalDataValues struct {
	Object string
	Price  int
}

// HistoricalTradeDataReturned type definition - used to define how trade data is returned
type HistoricalTradeDataReturned struct {
	Date time.Time
	HistoricalDataValues
}

// ErrTradeNotFound - error message for when a trade is not found
var ErrTradeNotFound = errors.New("trade not found")

// AlertType type definition - part of enum definition
type AlertType int

// Constants of alertType - 2nd part of enum definition
const (
	PriceAlertLowPrice  AlertType = iota // 0
	PriceAlertHighPrice                  // 1
)

// AlertValues type definition
type AlertValues struct {
	AlertType    AlertType
	PriceTrigger int
}

// AlertDef type definition
type AlertDef struct {
	Item string
	AlertValues
}

// HistoricalTradeAlertReturned - used to return the historical trades that match specific alerts
type HistoricalTradeAlertReturned struct {
	HistoricalTradeDataReturned
	AlertValues
}

// ErrNoDataMatchingAlertsFound - error message for when a trade is not found
var ErrNoDataMatchingAlertsFound = errors.New("no trades matching alerts found")
