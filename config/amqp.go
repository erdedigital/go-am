package config

import (
	"github.com/streadway/amqp"
)

func ConnAMQP() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		panic(err)
	}
	return conn
}