package store

import "Price_Notification_System/models"

type AlertDefStore interface {
	AddAlert(item string, newAlertDef models.AlertValues) error
	GetAlertsByItem(item string) ([]models.AlertsByItemReturned, error)
	GetAllAlerts() ([]models.AlertDef, error)
}
