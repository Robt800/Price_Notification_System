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

// LoadEnvVariablesPostgres function loads the environment variables from the .env file
func LoadEnvVariablesPostgres() (EnvVariablesPostgres, error) {

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

	// Check that the required environment variables are set, i.e. not nil
	if envVariables["POSTGRES_USER"] == "" || envVariables["POSTGRES_PASSWORD"] == "" || envVariables["POSTGRES_DB_NAME"] == "" {
		return EnvVariablesPostgres{"", "", ""}, fmt.Errorf("required environment variables are not set: POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB_NAME")
	}

	// Assuming no errors above - initialize the connection variables
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

// EnvVariablesBrevo holds the environment variables for Mailchimp connection
type EnvVariablesBrevo struct {
	BrevoAPIKey        string
	SenderEmailAddress string
}

// LoadEnvVariablesBrevo function loads the environment variables from the .env file
func LoadEnvVariablesBrevo() (EnvVariablesBrevo, error) {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		return EnvVariablesBrevo{"", ""}, fmt.Errorf("error loading the .env file: %v", err)
	}

	// Read the environment variables from the .env file
	envVariables, errReadingEnvVariables := godotenv.Read(".env")
	if errReadingEnvVariables != nil {
		return EnvVariablesBrevo{"", ""}, fmt.Errorf("error reading the .env file: %v", errReadingEnvVariables)
	}

	// Check that the required environment variables are set, i.e. not nil
	if (envVariables["BREVO_API_KEY"] == "") ||
		(envVariables["SENDER_EMAIL_ADDRESS"] == "") {
		return EnvVariablesBrevo{"", ""}, fmt.Errorf("required environment variables are not set: BREVO_API_KEY or SENDER_EMAIL_ADDRESS")
	}

	// Assuming no errors above - initialize the connection variables
	return EnvVariablesBrevo{
		BrevoAPIKey:        envVariables["BREVO_API_KEY"],
		SenderEmailAddress: envVariables["SENDER_EMAIL_ADDRESS"],
	}, nil
}
