package emails

import (
	"Price_Notification_System/models"
	"context"
)

type EmailClient interface {
	SendEmail(ctx context.Context, parameters models.EmailParameters) (status string, err error)
}
