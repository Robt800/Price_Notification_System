package service

import (
	"Price_Notification_System/producer/store"
)

// HistoricalStore - type related to the store.Store interface
type HistoricalStore struct {
	hist store.Storage
}

// NewHistoricalService - func to encapsulate store.Storage interface to the HistoricalStore local to this package
func NewHistoricalService(h store.Storage) HistoricalStore {
	return HistoricalStore{hist: h}
}

//func (h HistoricalStore) AlertRequired(al AlertStore) {

//}
