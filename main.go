package main

import (
	"fmt"

	"github.com/erdedigital/go-amqp/amqp"
)

const (
	routingKey string = "testing"
)

func main() {
	publiseh := amqp.TopicPublish()
	for i := 1; i < 5000; i++ {
		var payload = fmt.Sprintf("[%d] - %s", i, "OK")
		publiseh.Push(routingKey, payload)
	}
}
