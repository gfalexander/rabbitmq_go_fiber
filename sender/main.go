package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
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

	// Instance Queues
	_, err = channelRabbitMq.QueueDeclare("MessageService", true, false, false, false, nil)

	if err != nil {
		panic(err)
	}

	// Instance Fiber
	app := fiber.New()

	// Add Middleware
	app.Use(logger.New())

	// Create a route
	app.Get("/send", func(c *fiber.Ctx) error {
		// Create a message to publish
		message := amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(c.Query("message")),
		}

		if err := channelRabbitMq.Publish(
			"",
			"MessageService",
			false,
			false,
			message,
		); err != nil {
			return err
		}

		return nil
	})

	log.Fatal(app.Listen(":3000"))
}
