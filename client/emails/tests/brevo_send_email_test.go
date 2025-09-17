package tests

import (
	"Price_Notification_System/client/emails"
	"Price_Notification_System/config"
	"Price_Notification_System/models"
	"testing"
)

func TestBrevoSendEmail(t *testing.T) {
	//Load the environment variables
	enVariables, err := config.LoadEnvVariablesBrevo()
	if err != nil {
		t.Fatalf("Error loading environment variables: %v", err)
	}

	//Create an instance of the Brevo API
	brevoClient := emails.NewBrevoClient(enVariables.BrevoAPIKey)

	type testDefBrevoSendEmail struct {
		instanceName    string
		brevoClient     emails.EmailClient
		emailParameters models.EmailParameters
		expected        string
		wantErr         bool
	}

	tests := []testDefBrevoSendEmail{
		{instanceName: "test1 - Send email with valid parameters",
			brevoClient: brevoClient,
			emailParameters: models.EmailParameters{
				SenderEmail:    "rob@adsgb.co.uk",
				RecipientEmail: "rob@wealthresources.co.uk",
				RecipientName:  "Max",
				Subject:        "Test Email from Price Notification System",
				BodyText:       "<h1>This is a test email sent from the Price Notification System</h1><p>If you have received this email, the test has been successful.</p>",
			},
			expected: "202 Accepted",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.instanceName, func(t *testing.T) {
			status, err := tt.brevoClient.SendEmail(tt.emailParameters)
			if (err != nil) != tt.wantErr {
				t.Errorf("BrevoClient.SendEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if status != tt.expected {
				t.Errorf("BrevoClient.SendEmail() = %v, expected %v", status, tt.expected)
			}
		})
	}
}
