package emails

import "Price_Notification_System/models"

type EmailClient interface {
	SendEmail(parameters models.EmailParameters) (status string, err error)
}
