package store

import "Price_Notification_System/models"

type AlertDefStore interface {
	AddAlert(itemToAlert string, newAlertDef models.AlertValues, emailRecipient string) error
	GetAlertsByItem(item string) ([]models.AlertsByItemReturned, error)
	GetAllAlerts() ([]models.AlertDef, error)
}
