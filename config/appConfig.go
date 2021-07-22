package config

import (
	"crypto/tls"
	"net/http"

	"github.com/alexflint/go-arg"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type args struct {
	LogLevel           string `arg:"env:LOG_LEVEL"`
	Port               int    `arg:"env:PORT"`
	RabbitMqHost       string `arg:"env:RABBITMQ_HOST"`
	RabbitMqPort       string `arg:"env:RABBITMQ_PORT"`
	RabbitMqUser       string `arg:"env:RABBITMQ_USER"`
	RabbitMqPass       string `arg:"env:RABBITMQ_PASS"`
	FirebaseQueue      string `arg:"env:FIREBASE_NOTIFICATION_QUEUE"`
	FirebaseDLQ        string `arg:"env:FIREBASE_NOTIFICATION_DLQ"`
	FirebaseConfigPath string `arg:"env:FIREBASE_CONFIG_PATH"`
}

// Props is global variable for environment variables usage
var Props args

var opts struct {
	Profile string `arg:"-p" default:"default"`
}

// LoadConfig loads configuration when project run
func LoadConfig() {
	arg.MustParse(&opts)
	initLogger()
	log.Info("Application is starting with profile: ", opts.Profile)

	initEnvVars()
	_ = arg.Parse(&Props)
	applyLoggerLevel()
}

func initEnvVars() {
	if godotenv.Load("config/profiles/default.env") != nil {
		log.Fatal("Error in loading environment variables from: profiles/default.env")
	} else {
		log.Info("Environment variables loaded from: profiles/default.env")
	}

	if opts.Profile != "default" {
		profileFileName := "config/profiles/" + opts.Profile + ".env"
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		if godotenv.Overload(profileFileName) != nil {
			log.Fatal("Error in loading environment variables from: ", profileFileName)
		} else {
			log.Info("Environment variables overloaded from: ", profileFileName)
		}
	}
}

func initLogger() {
	log.SetLevel(log.InfoLevel)
	if opts.Profile == "default" {
		log.SetFormatter(&log.JSONFormatter{})
	}
}

func applyLoggerLevel() {
	loglevel, err := log.ParseLevel(Props.LogLevel)
	if err != nil {
		loglevel = log.InfoLevel
	}
	log.SetLevel(loglevel)
}
