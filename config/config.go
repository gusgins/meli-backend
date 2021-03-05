package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Configuration exported
type Configuration struct {
	API      APIConfiguration
	Database DatabaseConfiguration
}

// APIConfiguration exported
type APIConfiguration struct {
	APIPort int
}

// DatabaseConfiguration exported
type DatabaseConfiguration struct {
	DBHost     string
	DBPort     int
	DBName     string
	DBUser     string
	DBPassword string
}

// NewConfig returns system configuration
func NewConfig() Configuration {
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var configuration Configuration

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	// Set default value for variables
	viper.SetDefault("database.dbhost", "localhost")

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	return configuration
}
