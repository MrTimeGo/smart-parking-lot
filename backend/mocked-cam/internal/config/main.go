package config

import (
	rabbitmq "github.com/MrTimeGo/smart-parking-lot/backend/mocked-cam/internal/rabbitmq/config"
	s3 "github.com/MrTimeGo/smart-parking-lot/backend/mocked-cam/internal/s3/config"
	streamer "github.com/MrTimeGo/smart-parking-lot/backend/mocked-cam/internal/streamer/config"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type Config interface {
	comfig.Logger
	rabbitmq.AmqpConfigurator
	s3.S3Configurator
	streamer.StreamerConfigurator
}

type config struct {
	getter kv.Getter
	rabbitmq.AmqpConfigurator
	s3.S3Configurator
	streamer.StreamerConfigurator
	comfig.Logger
}

func New(getter kv.Getter) Config {
	return &config{
		getter:               getter,
		StreamerConfigurator: streamer.NewConfigurator(getter),
		S3Configurator:       s3.NewConfigurator(getter),
		AmqpConfigurator:     rabbitmq.NewConfigurator(getter),
		Logger:               comfig.NewLogger(getter, comfig.LoggerOpts{}),
	}
}
