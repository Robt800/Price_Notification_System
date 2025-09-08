package store

import (
	"Price_Notification_System/models"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// DBAlertStore - concrete implementation of the AlertDefStore interface
type DBAlertStore struct {
	db *sql.DB
}

// compile time check to ensure that inMemoryTradeStore implements the TradeStore interface
// This won't be used within the code - but if not implemented correctly - the compiler will throw an error
var _ AlertDefStore = &DBAlertStore{}

// NewDBAlertStore - creates a new instance of DBAlertStore
func NewDBAlertStore(dbConnStr string) (AlertDefStore, error) {

	// Initialize the DBAlertStore with a database connection
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	// Ping the database to verify connection
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to connect (via ping) to the database: %v", err)
	}

	//return &DBAlertStore{
	return &DBAlertStore{
		db: db,
	}, nil
}

// NewDBAlertStoreWithData - creates a new instance of DBAlertStore with pre-populated data
func NewDBAlertStoreWithData(dbConnStr string, data []models.AlertDef) (AlertDefStore, error) {

	// Initialize the DBAlertStore with a database connection
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	// Ping the database to verify connection
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	// Seed the database with initial data
	for _, alert := range data {
		// Replace with your actual insert logic and fields
		_, errDBExec := db.Exec(
			"INSERT INTO alerts (item, alert_type, price_trigger, email_recipient) VALUES ($1, $2, $3, $4)",
			alert.Item, alert.AlertType, alert.PriceTrigger, alert.EmailRecipient)
		if errDBExec != nil {
			return nil, fmt.Errorf("failed to insert values into the database: %v", err)
		}
	}

	// Return a new instance of DBAlertStore with the database connection
	return &DBAlertStore{
		db: db,
	}, nil
}

// AddAlert - adds a new alert to the alerts active - i.e. the private memory store used to facilitate easier testing
func (i *DBAlertStore) AddAlert(itemToAlert string, newAlertDef models.AlertValues, emailRecipient string) error {

	_, err := i.db.Exec("INSERT INTO alerts (item, alert_type, price_trigger, email_recipient) VALUES ($1, $2, $3, $4)",
		itemToAlert, newAlertDef.AlertType, newAlertDef.PriceTrigger, emailRecipient)
	if err != nil {
		return fmt.Errorf("failed to insert values into the database: %v", err)
	} else {
		fmt.Println("Data inserted successfully!")
	}

	return nil
}

// GetAlertsByItem - retrieves the alerts for a specific item
func (i *DBAlertStore) GetAlertsByItem(item string) (data []models.AlertsByItemReturned, err error) {

	//Testing - see value of item
	fmt.Printf("Querying for item: '%v'\n", item)

	// Query the database for all alerts
	rows, err := i.db.Query("SELECT id, item, alert_type, price_trigger, email_recipient FROM alerts WHERE item = $1", item)
	if err != nil {
		return nil, fmt.Errorf("failed to query data: %v", err)
	}

	for rows.Next() {
		var alert models.AlertDef
		var id int
		errRowsScan := rows.Scan(&id, &alert.Item, &alert.AlertType, &alert.PriceTrigger, &alert.EmailRecipient)
		if errRowsScan != nil {
			return nil, fmt.Errorf("failed to scan row: %v", errRowsScan)
		}

		data = append(data, models.AlertsByItemReturned{
			Item: alert.Item,
			AlertValues: models.AlertValues{
				AlertType:    alert.AlertType,
				PriceTrigger: alert.PriceTrigger},
			EmailRecipient: alert.EmailRecipient,
		})
	}

	// Close the rows after processing
	defer rows.Close()

	// Check if any data was found and execute the returns
	if len(data) > 0 {
		return data, nil
	} else {
		return nil, models.ErrNoAlertsForItemFound
	}
}

// GetAllAlerts - retrieves all the alerts
func (i *DBAlertStore) GetAllAlerts() (data []models.AlertDef, err error) {

	// Query the database for all alerts
	rows, err := i.db.Query("SELECT id, item, alert_type, price_trigger, email_recipient FROM alerts")
	if err != nil {
		return nil, fmt.Errorf("failed to query data: %v", err)
	}

	for rows.Next() {
		var alert models.AlertDef
		var id int
		errRowsScan := rows.Scan(&id, &alert.Item, &alert.AlertType, &alert.PriceTrigger, &alert.EmailRecipient)
		if errRowsScan != nil {
			return nil, fmt.Errorf("failed to scan row: %v", errRowsScan)
		}

		data = append(data, models.AlertDef{
			Item: alert.Item,
			AlertValues: models.AlertValues{
				AlertType:    alert.AlertType,
				PriceTrigger: alert.PriceTrigger},
			EmailRecipient: alert.EmailRecipient,
		})
	}

	if len(data) > 0 {
		return data, nil
	} else {
		return nil, fmt.Errorf("no alerts found")
	}
}
