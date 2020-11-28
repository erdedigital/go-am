package main

import (
	"log"
	"fmt"
	"os"


	"github.com/erdedigital/go-amqp/config"
	"github.com/erdedigital/go-amqp/exchange"
)

func main()  {
	log.Println("Hello")
	conn := config.ConnAMQP()

	emitter, err := exchange.NewEventEmitter(conn)
	if err != nil {
		panic(err)
	}

	for i := 1; i < 10; i++ {
		emitter.Push(fmt.Sprintf("[%d] - %s", i, os.Args[1]), os.Args[1])
	}
}	