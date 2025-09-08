package emails

import (
	"Price_Notification_System/models"
	"context"
	"fmt"
	brevo "github.com/getbrevo/brevo-go/lib"
)

type BrevoClient struct {
	api *brevo.APIClient
	ctx context.Context
}

// compile time check to ensure that inMemoryTradeStore implements the TradeStore interface
// This won't be used within the code - but if not implemented correctly - the compiler will throw an error
var _ EmailClient = &BrevoClient{}

// Constructor function to create a new instance of the Brevo Email Client
func NewBrevoClient(apiKey string) *BrevoClient {
	cfg := brevo.NewConfiguration()
	cfg.AddDefaultHeader("api-key", apiKey)
	client := brevo.NewAPIClient(cfg)
	return &BrevoClient{
		api: client,
		ctx: context.Background(),
	}
}

func (b *BrevoClient) SendEmail(parameters models.EmailParameters) (status string, err error) {
	//Build an instance of the SendSmtpEmail struct with the necessary parameters
	email := brevo.SendSmtpEmail{
		Sender: &brevo.SendSmtpEmailSender{
			Email: parameters.SenderEmail,
		},
		To: []brevo.SendSmtpEmailTo{
			{Email: parameters.RecipientEmail, Name: parameters.RecipientName},
		},
		Subject:     parameters.Subject,
		HtmlContent: parameters.BodyText,
	}

	//Make the API call to send the email
	_, resp, err := b.api.TransactionalEmailsApi.SendTransacEmail(b.ctx, email)
	if err != nil {
		fmt.Printf("Error when calling `TransactionalEmailsApi.SendTransacEmail`: %v\n", err)
		return "Failed", err
	}
	defer resp.Body.Close()

	return resp.Status, nil

}
