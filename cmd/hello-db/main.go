package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type alert struct {
	id           int
	item         string
	alertType    int
	priceTrigger int
}

func main() {
	// Replace with your actual credentials
	connStr := "user=postgres password=winter101 dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ping the database to verify connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to PostgreSQL successfully!")

	//Test Write Operation to existing table
	//_, err = db.Exec("INSERT INTO items (item) VALUES ($1)", "Batman")
	//if err != nil {
	//	log.Fatalf("Failed to insert data: %v", err)
	//} else {
	//	fmt.Println("Data inserted successfully!")
	//}

	_, err = db.Exec("INSERT INTO alerts (item, alert_type, price_trigger) VALUES ($1, $2, $3)", "Batman", 0, 500)
	if err != nil {
		log.Fatalf("Failed to insert data: %v", err)
	} else {
		fmt.Println("Data inserted successfully!")
	}

	//Test Read Operation from existing table
	rows, err := db.Query("SELECT id, item, alert_type, price_trigger FROM alerts")
	if err != nil {
		log.Fatalf("Failed to query data: %v", err)
	}

	defer rows.Close()

	var alerts []alert
	// Iterate through the result set
	for rows.Next() {
		var alert alert
		if err := rows.Scan(&alert.id, &alert.item, &alert.alertType, &alert.priceTrigger); err != nil {
			log.Fatalf("Failed to scan row: %v", err)
		}
		alerts = append(alerts, alert)
	}

	fmt.Printf("Alerts fetched from database: %+v\n", alerts)

}
