package config

import (
	"os"
)

//AppConfig APplication Configuration
var AppConfig Config

//Config Application configuration
type Config struct {
	Port string
}

// LoadConfig loads app configuration from env variables
func LoadConfig() (err error) {
	if os.Getenv("PORT") != "" {
		AppConfig.Port = os.Getenv("PORT")
	} else {
		AppConfig.Port = "8080"
	}
	return nil
}
