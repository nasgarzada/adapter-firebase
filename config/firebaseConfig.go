package config

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

type IFirebaseConfig interface {
	GetFirebaseMessaging() *messaging.Client
}

type FirebaseConfig struct {
}

func NewFirebaseConfig() IFirebaseConfig {
	return &FirebaseConfig{}
}

func (*FirebaseConfig) GetFirebaseMessaging() *messaging.Client {
	log.Info("ActionLog.GetFirebaseMessaging.start")
	opt := option.WithCredentialsFile(Props.FirebaseConfigPath)
	app, err := firebase.NewApp(context.Background(), &firebase.Config{}, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	firebaseMessaging, err := app.Messaging(context.Background())
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	log.Info("ActionLog.GetFirebaseMessaging.start")
	return firebaseMessaging
}
