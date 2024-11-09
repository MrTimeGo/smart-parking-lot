package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type Config interface {
	comfig.Logger
}

type config struct {
	getter kv.Getter
	comfig.Logger
}

func New(getter kv.Getter) Config {
	return &config{
		getter: getter,
		Logger: comfig.NewLogger(getter, comfig.LoggerOpts{}),
	}
}
