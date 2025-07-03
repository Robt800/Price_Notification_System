package store

import (
	"Price_Notification_System/models"
	"database/sql"
	"fmt"
	"log"
)

// DBAlertStore - concrete implementation of the AlertDefStore interface
type DBAlertStore struct {
	db *sql.DB
}

// compile time check to ensure that inMemoryTradeStore implements the TradeStore interface
// This won't be used within the code - but if not implemented correctly - the compiler will throw an error
var _ AlertDefStore = &DBAlertStore{}

// NewDBAlertStore - creates a new instance of DBAlertStore
func NewDBAlertStore(dbConnStr string) AlertDefStore {

	// Initialize the DBAlertStore with a database connection
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to verify connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	//return &DBAlertStore{
	return &DBAlertStore{
		db: db,
	}
}

// NewDBAlertStoreWithData - creates a new instance of DBAlertStore with pre-populated data
func NewDBAlertStoreWithData(dbConnStr string, data []models.AlertDef) AlertDefStore {

	// Initialize the DBAlertStore with a database connection
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to verify connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Seed the database with initial data
	for _, alert := range data {
		// Replace with your actual insert logic and fields
		_, errDBExec := db.Exec(
			"INSERT INTO alerts (item, alert_type, price_trigger) VALUES ($1, $2, $3)", alert.Item, alert.AlertType, alert.PriceTrigger,
		)
		if errDBExec != nil {
			log.Fatal("Failed to insert initial data: ", errDBExec)
		}
	}

	// Return a new instance of DBAlertStore with the database connection
	return &DBAlertStore{
		db: db,
	}
}

// AddAlert - adds a new alert to the alerts active - i.e. the private memory store used to facilitate easier testing
func (i *DBAlertStore) AddAlert(itemToAlert string, newAlertDef models.AlertValues) {

	_, err := i.db.Exec("INSERT INTO alerts (item, alert_type, price_trigger) VALUES ($1, $2, $3)", itemToAlert, newAlertDef.AlertType, newAlertDef.PriceTrigger)
	if err != nil {
		log.Fatalf("Failed to insert data: %v", err)
	} else {
		fmt.Println("Data inserted successfully!")
	}

}

// GetAlertsByItem - retrieves the alerts for a specific item
func (i *DBAlertStore) GetAlertsByItem(item string) (data []models.AlertsByItemReturned, err error) {

	// Query the database for all alerts
	rows, err := i.db.Query("SELECT id, item, alert_type, price_trigger FROM alerts")
	if err != nil {

		log.Fatalf("Failed to query data: %v", err)
	}

	for rows.Next() {
		var alert models.AlertDef
		errRowsScan := rows.Scan(&alert.Item, &alert.AlertType, &alert.PriceTrigger)
		if errRowsScan != nil {
			log.Fatalf("Failed to scan row: %v", errRowsScan)
		}

		if alert.Item == item {
			data = append(data, models.AlertsByItemReturned{
				Item: item,
				AlertValues: models.AlertValues{
					AlertType:    alert.AlertType,
					PriceTrigger: alert.PriceTrigger},
			})
		}
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
	rows, err := i.db.Query("SELECT id, item, alert_type, price_trigger FROM alerts")
	if err != nil {

		log.Fatalf("Failed to query data: %v", err)
	}

	for rows.Next() {
		var alert models.AlertDef
		errRowsScan := rows.Scan(&alert.Item, &alert.AlertType, &alert.PriceTrigger)
		if errRowsScan != nil {
			log.Fatalf("Failed to scan row: %v", errRowsScan)
		}

		data = append(data, models.AlertDef{
			Item: alert.Item,
			AlertValues: models.AlertValues{
				AlertType:    alert.AlertType,
				PriceTrigger: alert.PriceTrigger},
		})
	}

	if len(data) > 0 {
		return data, nil
	} else {
		return nil, fmt.Errorf("no alerts found")
	}
}
