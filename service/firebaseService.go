package service

import (
	"context"
	"encoding/json"
	"firebase.google.com/go/v4/messaging"
	"github.com/nasgarzada/adapter-firebase/config"
	"github.com/nasgarzada/adapter-firebase/model"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"reflect"
)

type IFirebaseService interface {
	HandleNotificationQueue(msg amqp.Delivery) error
}

type FirebaseServiceImpl struct {
	FirebaseConfig config.IFirebaseConfig
}

func NewFirebaseService(firebaseConfig config.IFirebaseConfig) IFirebaseService {
	return &FirebaseServiceImpl{
		FirebaseConfig: firebaseConfig,
	}
}

func (f *FirebaseServiceImpl) SendNotification(notification *model.Notification) {
	log.Infof("ActionLog.SendNotification.start - title:%s", notification.Title)
	firebaseMessaging := f.FirebaseConfig.GetFirebaseMessaging()
	multicastMessage := &messaging.MulticastMessage{
		Tokens: notification.DeviceToken,
		Notification: &messaging.Notification{
			Title: notification.Title,
			Body:  notification.Body,
		},
	}
	if notification.Data != nil && notification.Data.Payload != nil {
		multicastMessage.Data = notification.Data.Payload

		notificationGroup := reflect.TypeOf(notification.Data.NotificationGroup).Name()
		notificationType := reflect.TypeOf(notification.Data.NotificationType).Name()

		multicastMessage.Data[notificationGroup] = string(notification.Data.NotificationGroup)
		multicastMessage.Data[notificationType] = string(notification.Data.NotificationType)

		if notification.Data.ImageUrl != nil {
			multicastMessage.Data[reflect.TypeOf(notification.Data.ImageUrl).Name()] = *notification.Data.ImageUrl
		}
	}

	batchResponse, err := firebaseMessaging.SendMulticast(context.Background(), multicastMessage)
	if err != nil {
		log.Fatalln(err)
	}

	for _, v := range batchResponse.Responses {
		log.Debugf("Successful sends : %v", v)
	}

	log.Infof("ActionLog.SendNotification.end - title:%s", notification.Title)
}

func (f *FirebaseServiceImpl) HandleNotificationQueue(msg amqp.Delivery) error {
	log.Info("ActionLog.HandleNotificationQueue.start")
	notification := new(model.Notification)
	err := json.Unmarshal(msg.Body, &notification)
	if err != nil {
		log.Errorf("error happened on HandleNotificationQueue %v", err)
		return err
	}
	f.SendNotification(notification)

	log.Info("ActionLog.HandleNotificationQueue.end")
	return nil
}
