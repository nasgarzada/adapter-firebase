package main

import (
	"github.com/go-chi/chi"
	"github.com/nasgarzada/adapter-firebase/config"
	"github.com/nasgarzada/adapter-firebase/handler"
	"github.com/nasgarzada/adapter-firebase/queue"
	"github.com/nasgarzada/adapter-firebase/service"
	"github.com/nasgarzada/adapter-firebase/util"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"net/http"
	"strconv"
)

func main() {
	config.LoadConfig()

	if err := util.CreateFirebaseConfigFile(); err != nil {
		panic("Couldn't create config file")
	}

	router := chi.NewRouter()

	handler.NewHealthHandler(router)

	firebaseConfig := config.NewFirebaseConfig()
	firebaseService := service.NewFirebaseService(firebaseConfig)

	go runRatingSurveyQueueListener(firebaseService.HandleNotificationQueue)

	port := strconv.Itoa(config.Props.Port)
	log.Info("Starting server at port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func runRatingSurveyQueueListener(listener func(msg amqp.Delivery) error) {
	queue.ReceiveMessages(listener)
}
