package queue

import (
	"fmt"
	"github.com/nasgarzada/adapter-firebase/config"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func ReceiveMessages(listener func(msg amqp.Delivery) error) {
	receiveMessage(listener, config.Props.FirebaseQueue, config.Props.FirebaseDLQ)
}

func receiveMessage(listener func(msg amqp.Delivery) error, queueName, dlQueueName string) {
	log.Infof("ActionLog.receiveMessage.start - queue name: %s, dlq: %s", queueName, dlQueueName)
	ch, conn := newRabbitChannel()
	defer ch.Close()
	defer conn.Close()

	dlqExchangeName := fmt.Sprintf("%s_Exchange", dlQueueName)
	dlqKeyName := fmt.Sprintf("%s_Key", dlQueueName)

	err := ch.ExchangeDeclare(
		dlqExchangeName,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)

	failOnError(err, fmt.Sprintf("Failed to declare an exchange %s", dlqExchangeName))

	args := amqp.Table{
		"x-dead-letter-exchange":    dlqExchangeName,
		"x-dead-letter-routing-key": dlqKeyName,
		"x-message-ttl":             300000,
	}

	declaredQueue, err := ch.QueueDeclare(queueName, true, false, false, false, args)
	failOnError(err, fmt.Sprintf("Failed to declare queue %s", queueName))

	declaredDlQueue, err := ch.QueueDeclare(dlQueueName, true, false, false, false, nil)
	failOnError(err, fmt.Sprintf("Failed to declare queue %s", dlQueueName))

	err = ch.QueueBind(
		declaredDlQueue.Name,
		dlqKeyName,
		dlqExchangeName,
		false,
		nil,
	)
	failOnError(err, fmt.Sprintf("Failed to bind queue %s", queueName))

	messages, err := ch.Consume(
		declaredQueue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, fmt.Sprintf("Failed to register a consumer %s", queueName))
	forever := make(chan bool)
	go func() {
		for msg := range messages {
			err = listener(msg)
			if err != nil {
				_ = msg.Nack(false, false)
				continue
			}
			_ = msg.Ack(false)
		}
	}()
	log.Printf("Waiting for messages.")
	<-forever

	log.Infof("ActionLog.receiveMessage.end - queue name: %s, dlq: %s", queueName, dlQueueName)
}
