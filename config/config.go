package config

import (
	"github.com/joho/godotenv"
	"log"
)

// EnvVariablesPostgres holds the environment variables for PostgreSQL connection
type EnvVariablesPostgres struct {
	PostgresUser     string
	PostgresPassword string
	PostgresDBName   string
}

// ConnVarPostgres is a global variable that holds the PostgreSQL connection variables
var ConnVarPostgres EnvVariablesPostgres

// DBConnStr is the connection string for the PostgreSQL database
var DBConnStr string

// init function initialises the environment variables from the .env file
func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

	// Read the environment variables from the .env file
	envVariables, errReadingEnvVariables := godotenv.Read(".env")
	if errReadingEnvVariables != nil {
		log.Fatalf("Error reading .env file: %v", errReadingEnvVariables)
	}

	// Initialize the connection variables
	ConnVarPostgres.PostgresUser = envVariables["POSTGRES_USER"]
	ConnVarPostgres.PostgresPassword = envVariables["POSTGRES_PASSWORD"]
	ConnVarPostgres.PostgresDBName = envVariables["POSTGRES_DB_NAME"]
}

// init function constructs the PostgreSQL connection string
func init() {

	DBConnStr = "user=" + ConnVarPostgres.PostgresUser +
		" password=" + ConnVarPostgres.PostgresPassword +
		" dbname=" + ConnVarPostgres.PostgresDBName +
		" sslmode=disable"

	log.Println("PostgreSQL connection string initialized successfully.")
}
