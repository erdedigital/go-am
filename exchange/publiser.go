package exchange

import (
	"log"
	"github.com/erdedigital/go-amqp/exechange/services"
	"github.com/streadway/amqp"
)

// Emitter for publishing AMQP events
type Emitter struct {
	connection *amqp.Connection
}

func (e *Emitter) setup() error {
	channel, err := e.connection.Channel()
	if err != nil {
		panic(err)
	}

	defer channel.Close()
	return services.declareExchange(channel)
}

// Push (Publish) a specified message to the AMQP exchange
func (e *Emitter) Push(event string, severity string) error {
	channel, err := e.connection.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()

	err = channel.Publish(
		services.getExchangeName(),
		severity,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(event),
		},
	)
	log.Printf("Sending message: %s -> %s", event, services.getExchangeName())
	return nil
}

// NewEventEmitter returns a new event.Emitter object
// ensuring that the object is initialised, without error
func NewEventEmitter(conn *amqp.Connection) (Emitter, error) {
	emitter := Emitter{
		connection: conn,
	}

	err := emitter.setup()
	if err != nil {
		return Emitter{}, err
	}

	return emitter, nil
}