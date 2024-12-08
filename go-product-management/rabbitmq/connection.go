package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
)

var Conn *amqp.Connection
var Channel *amqp.Channel

func InitRabbitMQ() {
	var err error
	Conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	Channel, err = Conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	log.Println("RabbitMQ connection established.")
}
