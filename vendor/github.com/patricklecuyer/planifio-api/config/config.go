package config

import (
	"os"

	"github.com/patricklecuyer/planifio-api/datastore"
)

//AppConfig APplication Configuration
var AppConfig Config

//Config Application configuration
type Config struct {
	Port            string
	MongoURI        string
	MongoDatabase   string
	RecaptchaSecret string
	Datastore       *datastore.Datastore
	JWTSecret       []byte
	Env             string
	RedisURL        string
}

// LoadConfig loads app configuration from env variables
func LoadConfig() (err error) {
	AppConfig.Port = os.Getenv("PORT")
	AppConfig.MongoURI = os.Getenv("MONGODB_URI")
	AppConfig.MongoDatabase = os.Getenv("MOGODB_DB")
	AppConfig.RecaptchaSecret = os.Getenv("RECPATCHA_SECRET")
	AppConfig.JWTSecret = []byte(os.Getenv("JWT_SECRET"))
	AppConfig.Env = os.Getenv("ENV")
	AppConfig.RedisURL = os.Getenv("REDIS_URL")
	AppConfig.Datastore, err = datastore.New(AppConfig.MongoURI, AppConfig.MongoDatabase)

	return
}
