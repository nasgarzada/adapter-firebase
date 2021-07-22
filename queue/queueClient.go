package queue

import (
	"fmt"
	"github.com/nasgarzada/adapter-firebase/config"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"os"
	"os/signal"
	"syscall"
)

func newRabbitChannel() (*amqp.Channel, *amqp.Connection) {
	rabbitUrl := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		config.Props.RabbitMqUser, config.Props.RabbitMqPass, config.Props.RabbitMqHost, config.Props.RabbitMqPort)

	conn, err := amqp.Dial(rabbitUrl)
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open channel")

	return ch, conn
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		gracefulShutdown()
	}
}

func gracefulShutdown() {
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	signal.Notify(s, syscall.SIGTERM)
	go func() {
		<-s
		log.Println("Shutting down gracefully.")
		// clean up here
		os.Exit(0)
	}()
}
