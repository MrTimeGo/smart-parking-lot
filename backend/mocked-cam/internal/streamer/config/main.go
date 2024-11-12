package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"time"
)

const cfgKey = "streamer"

type StreamOptions struct {
	MinDelay  time.Duration `fig:"min_delay,required"`
	MaxDelay  time.Duration `fig:"max_delay,required"`
	InitDelay time.Duration `fig:"init_delay,required"`
}

type StreamerConfig struct {
	DumpFile       string        `fig:"dump_file,required"`
	EnteredOptions StreamOptions `fig:"entered,required"`
	ExitedOptions  StreamOptions `fig:"exited,required"`
}

type StreamerConfigurator interface {
	StreamerConfig() StreamerConfig
}

type configurator struct {
	getter kv.Getter
	once   comfig.Once
}

func NewConfigurator(getter kv.Getter) StreamerConfigurator {
	return &configurator{getter: getter}
}

func (c *configurator) StreamerConfig() StreamerConfig {
	return c.once.Do(func() interface{} {
		var cfg StreamerConfig

		if figure.
			Out(&cfg).
			From(kv.MustGetStringMap(c.getter, cfgKey)).
			Please() != nil {
			panic("failed to load streamer config")
		}

		return cfg
	}).(StreamerConfig)
}
