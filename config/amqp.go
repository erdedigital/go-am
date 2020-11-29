package config

import (
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

func ConnAMQP() *amqp.Connection {
	amqpURL := fmt.Sprintf("amqp://%s", os.Getenv("AMQP_URL"))
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		panic(err)
	}
	return conn
}
