package rabbitmq

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)

type Message struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

func Publish(queueName string, message Message) {
	body, err := json.Marshal(message)
	if err != nil {
		log.Printf("Failed to encode message: %v", err)
		return
	}

	err = Channel.Publish(
		"",        // Exchange
		queueName, // Queue name
		false,     // Mandatory
		false,     // Immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Printf("Failed to publish message: %v", err)
	}
}
