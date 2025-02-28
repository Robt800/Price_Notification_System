package store

import (
	"time"
)

// OverallStore type that can hold either a Storage or Notification implementation
type OverallStore interface {
	Storage
	Notification
}

// overallStoreImpl - Concrete implementation of OverallStore
type overallStoreImpl struct {
	storage      Storage
	notification Notification
}

// Add - Implement the Storage interface
func (sw overallStoreImpl) Add(t time.Time, hdv HistoricalDataValues) {
	if sw.storage != nil {
		sw.storage.Add(t, hdv)
	}
}

// AddAlert - Implement the Notification interface
func (sw overallStoreImpl) AddAlert(item string, newAlertDef alert) {
	if sw.notification != nil {
		sw.notification.AddAlert(item, newAlertDef)
	}
}

// NewStoreFromStorage - function to create a OverallStore from a Storage implementation
func NewStoreFromStorage(s Storage) OverallStore {
	return overallStoreImpl{storage: s}
}

// NewStoreFromNotification - function to create a OverallStore from a Notification implementation
func NewStoreFromNotification(n Notification) OverallStore {
	return overallStoreImpl{notification: n}
}
