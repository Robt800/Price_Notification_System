package store

import (
	"Price_Notification_System/models"
	"time"
)

type TradeStore interface {
	AddTrade(time.Time, models.HistoricalDataValues) error
	GetTradeByItem(string) (data []models.HistoricalTradeDataReturned, err error)
}
