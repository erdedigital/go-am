package exchange

import (
	"log"

	"github.com/streadway/amqp"
)

// Publish for publishing AMQP events
type Publish struct {
	connection *amqp.Connection
}

func (e *Publish) setup() error {
	channel, err := e.connection.Channel()
	if err != nil {
		panic(err)
	}

	defer channel.Close()
	return declareExchange(channel)
}

// Push (Publish) a specified message to the AMQP exchange
func (e *Publish) Push(event string, routingKey string) error {
	channel, err := e.connection.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()

	err = channel.Publish(
		getExchangeName(),
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(event),
		},
	)
	log.Printf("Sending message: %s -> %s", event, getExchangeName())
	return nil
}

// NewEventPublish returns a new event.Publish object
// ensuring that the object is initialised, without error
func NewEventPublish(conn *amqp.Connection) (Publish, error) {
	publiser := Publish{
		connection: conn,
	}

	err := publiser.setup()
	if err != nil {
		return Publish{}, err
	}

	return publiser, nil
}
