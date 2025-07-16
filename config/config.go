package config

import (
	"fmt"
	"github.com/joho/godotenv"
)

// EnvVariablesPostgres holds the environment variables for PostgreSQL connection
type EnvVariablesPostgres struct {
	PostgresUser     string
	PostgresPassword string
	PostgresDBName   string
}

// LoadEnvVariables function loads the environment variables from the .env file
func LoadEnvVariables() (EnvVariablesPostgres, error) {

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		return EnvVariablesPostgres{"", "", ""}, fmt.Errorf("error loading the .env file: %v", err)
	}

	// Read the environment variables from the .env file
	envVariables, errReadingEnvVariables := godotenv.Read(".env")
	if errReadingEnvVariables != nil {
		return EnvVariablesPostgres{"", "", ""}, fmt.Errorf("error reading the .env file: %v", errReadingEnvVariables)
	}

	// Initialize the connection variables
	return EnvVariablesPostgres{
		PostgresUser:     envVariables["POSTGRES_USER"],
		PostgresPassword: envVariables["POSTGRES_PASSWORD"],
		PostgresDBName:   envVariables["POSTGRES_DB_NAME"],
	}, nil
}

// LoadDBConnectionStr function constructs the PostgreSQL connection string
func LoadDBConnectionStr(connVariables EnvVariablesPostgres) string {

	return fmt.Sprintf(
		"user=" + connVariables.PostgresUser +
			" password=" + connVariables.PostgresPassword +
			" dbname=" + connVariables.PostgresDBName +
			" sslmode=disable",
	)
}
