package store

import (
	"Price_Notification_System/models"
	"database/sql"
	"fmt"
	"time"
)

// DBTradeStore - concrete implementation of the TradeStore interface
type DBTradeStore struct {
	db *sql.DB
}

// compile time check to ensure that inMemoryTradeStore implements the TradeStore interface
// This won't be used within the code - but if not implemented correctly - the compiler will throw an error
var _ TradeStore = &DBTradeStore{}

// NewDBAlertStore - creates a new instance of DBAlertStore
func NewDBTradeStore(dbConnStr string) (TradeStore, error) {

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

	//return &DBAlertStore{
	return &DBTradeStore{
		db: db,
	}, nil
}

func (i *DBTradeStore) AddTrade(tradeTime time.Time, tradeValues models.HistoricalDataValues) error {

	_, err := i.db.Exec("INSERT INTO trades (time_stamp, item, price) VALUES ($1, $2, $3)", tradeTime, tradeValues.Object, tradeValues.Price)
	if err != nil {
		return fmt.Errorf("failed to insert values into the database: %v", err)
	} else {
		fmt.Println("Data inserted successfully!")
	}

	return nil
}

func (i *DBTradeStore) GetTradeByItem(item string) (data []models.HistoricalTradeDataReturned, err error) {

	// Query the database for all trades for the specified item
	rows, err := i.db.Query("SELECT time_stamp, item, price FROM trades WHERE item = $1", item)
	if err != nil {
		return nil, fmt.Errorf("failed to query data: %v", err)
	}

	for rows.Next() {
		var trade models.HistoricalTradeDataReturned
		errRowsScan := rows.Scan(&trade.Date, &trade.Object, &trade.Price)
		if errRowsScan != nil {
			return nil, fmt.Errorf("failed to scan row: %v", errRowsScan)
		}

		data = append(data, models.HistoricalTradeDataReturned{
			Date: trade.Date,
			HistoricalDataValues: models.HistoricalDataValues{
				Object: trade.Object,
				Price:  trade.Price},
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
