package store

import "Price_Notification_System/models"

type AlertDefStore interface {
	AddAlert(item string, newAlertDef models.AlertValues)
	ProcessAlerts(alertToProcess models.AlertDef) (data []models.HistoricalTradeAlertReturned, err error)
}
