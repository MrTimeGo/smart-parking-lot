package config

import (
	"github.com/pkg/errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"reflect"
)

const (
	cfgKey = "rabbitmq"
)

type AmqpConfig struct {
	Connection          *amqp.Connection `fig:"url,required"`
	EnteredCarQueueName string           `fig:"entered_car_queue_name,required"`
	ExitedCarQueueName  string           `fig:"exited_car_queue_name,required"`
}

type AmqpConfigurator interface {
	AmqpConfig() AmqpConfig
}

type configurator struct {
	getter kv.Getter
	once   comfig.Once
}

func NewConfigurator(getter kv.Getter) AmqpConfigurator {
	return &configurator{getter: getter}
}

func (c *configurator) AmqpConfig() AmqpConfig {
	return c.once.Do(func() interface{} {
		var cfg AmqpConfig

		if figure.
			Out(&cfg).
			With(figure.BaseHooks, rabbitHooks).
			From(kv.MustGetStringMap(c.getter, cfgKey)).
			Please() != nil {
			panic("failed to load rabbitmq config")
		}

		return cfg
	}).(AmqpConfig)
}

var rabbitHooks = figure.Hooks{
	"*amqp091.Connection": func(value interface{}) (reflect.Value, error) {
		switch v := value.(type) {
		case string:
			conn, err := amqp.Dial(v)
			if err != nil {
				return reflect.Value{}, errors.Wrap(err, "failed to dial rabbitmq")
			}

			return reflect.ValueOf(conn), nil
		default:
			return reflect.Value{}, errors.Errorf("failed to cast %#v of type %T to *rabbitmq.Connection", value, value)
		}
	},
}
