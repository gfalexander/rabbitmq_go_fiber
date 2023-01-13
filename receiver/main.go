package main

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

func main() {
	// Create a new RabbitMQ Connection
	connectionRabbitMq, err := amqp.Dial(os.Getenv("AMQP_SERVER_URL"))
	if err != nil {
		panic(err)
	}
	defer connectionRabbitMq.Close()

	// Create a new RabbitMQ Channel
	channelRabbitMq, err := connectionRabbitMq.Channel()
	if err != nil {
		panic(err)
	}
	defer channelRabbitMq.Close()

	// Subscribe in channel to get messages
	messages, err := channelRabbitMq.Consume("MessageService", "", true, false, false, false, nil)

	if err != nil {
		panic(err)
	}

	log.Println("Successfuly connected to RabbitMQ")
	log.Println("Waiting messages")

	forever := make(chan bool)

	go func() {
		for message := range messages {
			log.Printf("> Received message: %s\n", message.Body)
		}
	}()

	<-forever
}
