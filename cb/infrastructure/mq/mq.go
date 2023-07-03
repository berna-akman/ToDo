package mq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func InitMQ() *amqp.Channel {
	rabbitMQURL := "amqp://berna1:akman1@localhost:5672/"
	virtualHost := "to-do-api"

	conn, err := amqp.Dial(fmt.Sprintf("%s%s", rabbitMQURL, virtualHost))
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	//defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	//defer ch.Close()

	return ch
}

func ProduceMessage() error {
	queueName := "card_assignments"
	err := InitMQ().Publish("", queueName, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        []byte("Card Assigned"),
	})
	return err
}
