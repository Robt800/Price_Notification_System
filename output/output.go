package output

import (
	"Price_Notification_System/client/emails"
	"Price_Notification_System/models"
	"Price_Notification_System/producer/trades"
	store "Price_Notification_System/store"
	"context"
	"fmt"
	"io"
)

// Outputs ensures the data from the channel (i.e. the trade) is genuine - if it is, it prints it
func Outputs(ctx context.Context, producedData chan trades.TradeItems, tradeStore store.TradeStore, alertStore store.AlertDefStore, emailClient emails.EmailClient, emailSenderAddress string, write io.Writer) error {

	for {
		select {

		case <-ctx.Done():
			fmt.Fprintf(write, "Context cancelled\n")
			fmt.Fprintf(write, "The error was: %v\n", ctx.Err())
			return ctx.Err()

		case tradeData, ok := <-producedData:
			if !ok {
				return nil
			}
			_, _ = fmt.Fprintf(write, "%v\n", tradeData)

			// Check if the trade meets any alert criteria
			alerts, err := alertStore.GetAlertsByItem(tradeData.Object)
			if err != nil {
				fmt.Fprintf(write, "Error retrieving alerts for item %s: %v\n", tradeData.Object, err)
			}

			// Determine if the trade matches any alert criteria
			tradeMatchesAlertCriteria, mapAlertsByItem := tradeMeetsAlertCriteria(tradeData, alerts)

			if tradeMatchesAlertCriteria {
				errAddingTrade := tradeStore.AddTrade(tradeData.Timestamp, models.HistoricalDataValues{Object: tradeData.Object, Price: tradeData.Price})
				if errAddingTrade != nil {
					fmt.Fprintf(write, "Error storing trade data: %v\n", errAddingTrade)
				} else {
					fmt.Fprintf(write, "Trade data stored successfully!\n")
				}
				// Send email notification
				for _, v := range mapAlertsByItem {
					_, errSendingEmail := emailClient.SendEmail(ctx, models.EmailParameters{
						SenderEmail:    emailSenderAddress,
						RecipientEmail: v.EmailRecipient,
						RecipientName:  "",
						Subject:        fmt.Sprintf("Price Alert for %s", tradeData.Object),
						BodyText:       fmt.Sprintf("The price for %s has reached %d", tradeData.Object, tradeData.Price),
					})
					if errSendingEmail != nil {
						fmt.Fprintf(write, "Error sending email: %v\n", errSendingEmail)
					} else {
						fmt.Fprintf(write, "Email sent successfully to %s!\n", v.EmailRecipient)
					}
				}
			}
		}
	}
}

func tradeMeetsAlertCriteria(tradeData trades.TradeItems, alerts []models.AlertsByItemReturned) (match bool, alertsMatchingTrade []models.AlertsByItemReturned) {

	for _, alert := range alerts {

		if tradeData.Object == alert.Item &&
			alert.AlertType == 0 &&
			tradeData.Price <= alert.PriceTrigger {
			match = true
			alertsMatchingTrade = append(alertsMatchingTrade, alert)
		} else if tradeData.Object == alert.Item &&
			alert.AlertType == 1 &&
			tradeData.Price >= alert.PriceTrigger {
			match = true
			alertsMatchingTrade = append(alertsMatchingTrade, alert)
		}
	}
	return match, alertsMatchingTrade
}
