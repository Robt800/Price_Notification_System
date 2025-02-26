package service

import (
	"Price_Notification_System/producer/store"
)

// HistoricalStore - type related to the store.Storage interface
type HistoricalStore struct {
	hist store.Storage
}

// AlertStore - type related to the store.Notification
type AlertStore struct {
	notify store.Notification
}

// NewAlertService - func to encapsulate store.Storage interface to the HistoricalStore local to this package
func NewAlertService(h store.Storage, n store.Notification) (HistoricalStore, AlertStore) {
	return HistoricalStore{hist: h}, AlertStore{notify: n}
}

func (h HistoricalStore) AlertRequired(al AlertStore) {

}
