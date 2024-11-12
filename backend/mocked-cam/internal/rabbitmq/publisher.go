package rabbitmq

import (
	"github.com/MrTimeGo/smart-parking-lot/backend/mocked-cam/internal/rabbitmq/config"
	"github.com/pkg/errors"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	channel *amqp.Channel
	enterQ  string
	exitQ   string
}

func New(cfg config.AmqpConfig) *Publisher {
	ch, err := cfg.Connection.Channel()
	if err != nil {
		panic(errors.Wrap(err, "failed to create channel"))
	}

	_, err = ch.QueueDeclare(cfg.EnteredCarQueueName, false, false, false, false, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to declare entered car queue"))
	}

	_, err = ch.QueueDeclare(cfg.ExitedCarQueueName, false, false, false, false, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to declare exited car queue"))
	}

	return &Publisher{
		channel: ch,
		exitQ:   cfg.ExitedCarQueueName,
		enterQ:  cfg.EnteredCarQueueName,
	}
}

func (p *Publisher) PublishCar(queue, car string) error {
	return p.channel.Publish("", queue, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(car),
	})
}

func (p *Publisher) EnterQueueName() string {
	return p.enterQ
}

func (p *Publisher) ExitQueueName() string {
	return p.exitQ
}
