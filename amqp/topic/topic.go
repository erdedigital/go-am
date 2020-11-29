package topic

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// Publish for publishing AMQP events
type Publish struct {
	connection *amqp.Connection
}

var (
	exchangeName string = "webhook"
	typeName     string = "topic"
)

// NewEventPublish returns a new event.Publish object
// ensuring that the object is initialised, without error
func NewEventPublish(conn *amqp.Connection) (Publish, error) {
	publiser := Publish{
		connection: conn,
	}

	err := publiser.OpenChannel()
	if err != nil {
		return Publish{}, err
	}

	return publiser, nil
}

func declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",    // name
		true,  // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
}

func declareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		exchangeName, // name
		typeName,     // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
}

func (e *Publish) OpenChannel() error {
	channel, err := e.connection.Channel()
	if err != nil {
		panic(err)
	}

	defer channel.Close()
	return declareExchange(channel)
}

// Push (Publish) a specified message to the AMQP exchange
func (e *Publish) Push(routingKey string, payload string) error {
	var mainRoutingKey = fmt.Sprintf("%s.%s", exchangeName, routingKey)
	channel, err := e.connection.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()

	err = channel.Publish(
		exchangeName,
		mainRoutingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(payload),
		},
	)
	log.Printf("Sending message: %s -> %s -> %s", payload, exchangeName, mainRoutingKey)
	return nil
}
