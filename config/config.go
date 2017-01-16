package config

import (
	"os"

	"github.com/patricklecuyer/cr460-lab1/datastore"
)

//AppConfig APplication Configuration
var AppConfig Config

//Config Application configuration
type Config struct {
	Port          string
	MongoURI      string
	MongoDatabase string
	Datastore     *datastore.Datastore
}

// LoadConfig loads app configuration from env variables
func LoadConfig() (err error) {
	if os.Getenv("PORT") != "" {
		AppConfig.Port = os.Getenv("PORT")
	} else {
		AppConfig.Port = "8080"
	}

	AppConfig.MongoURI = os.Getenv("MONGODB_URI")
	AppConfig.MongoDatabase = os.Getenv("MOGODB_DB")
	AppConfig.Datastore, err = datastore.New(AppConfig.MongoURI, AppConfig.MongoDatabase)

	return
}
