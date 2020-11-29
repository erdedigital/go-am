package amqp

import (
	"github.com/erdedigital/go-amqp/amqp/topic"
	"github.com/erdedigital/go-amqp/config"
)

func TopicPublish() topic.Publish {
	conn := config.ConnAMQP()
	publiseh, err := topic.NewEventPublish(conn)
	if err != nil {
		panic(err)
	}

	return publiseh
}
