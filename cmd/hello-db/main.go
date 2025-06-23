package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
)

type alert struct {
	id           int
	item         string
	alertType    int
	priceTrigger int
}

type env struct {
	postgresUser     string
	postgresPassword string
	postgresDBName   string
}

func main() {

	//load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Read the environment variables from the .env file
	envVariables, errReadingEnvVariables := godotenv.Read(".env")
	if errReadingEnvVariables != nil {
		log.Fatalf("Error reading .env file: %v", errReadingEnvVariables)
	}

	// Print the environment variables to verify they are loaded correctly
	fmt.Printf("Postgres User: %s, Password: %s, DB Name: %s\n", envVariables["POSTGRES_USER"], envVariables["POSTGRES_PASSWORD"], envVariables["POSTGRES_DB_NAME"])

	// Replace with your actual credentials
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", envVariables["POSTGRES_USER"], envVariables["POSTGRES_PASSWORD"], envVariables["POSTGRES_DB_NAME"])

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
		var rowItem alert
		if err := rows.Scan(&rowItem.id, &rowItem.item, &rowItem.alertType, &rowItem.priceTrigger); err != nil {
			log.Fatalf("Failed to scan row: %v", err)
		}
		alerts = append(alerts, rowItem)
	}

	fmt.Printf("Alerts fetched from database: %+v\n", alerts)

}
