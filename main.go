package main

import (
	"fmt"
	"log"

	"github.com/erdedigital/go-amqp/config"
	"github.com/erdedigital/go-amqp/exchange"
)

func main() {
	log.Println("Hello")
	conn := config.ConnAMQP()

	publiseh, err := exchange.NewEventPublish(conn)
	if err != nil {
		panic(err)
	}

	for i := 1; i < 1000; i++ {
		publiseh.Push(fmt.Sprintf("[%d] - %s", i, "OK"), "testing")
	}
}
