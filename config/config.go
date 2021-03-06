package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// Configuration exported
type Configuration struct {
	API      APIConfiguration
	Database DatabaseConfiguration
}

// APIConfiguration exported
type APIConfiguration struct {
	Port int
}

// DatabaseConfiguration exported
type DatabaseConfiguration struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

// NewConfig returns system configuration
func NewConfig() Configuration {

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

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
