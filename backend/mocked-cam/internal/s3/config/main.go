package config

import (
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

const cfgKey = "s3"

type S3Config struct {
	Credentials S3Credentials `fig:"credentials,required"`
	Bucket      string        `fig:"bucket,required"`
}

type S3Credentials struct {
	Endpoint  string `fig:"endpoint,required"`
	AccessKey string `fig:"access_key,required"`
	SecretKey string `fig:"secret_key,required"`
}

type S3Configurator interface {
	S3Config() S3Config
}

type configurator struct {
	getter kv.Getter
	once   comfig.Once
}

func NewConfigurator(getter kv.Getter) S3Configurator {
	return &configurator{getter: getter}
}

func (c *configurator) S3Config() S3Config {
	return c.once.Do(func() interface{} {
		var cfg S3Config

		if err := figure.
			Out(&cfg).
			From(kv.MustGetStringMap(c.getter, cfgKey)).
			Please(); err != nil {
			panic("failed to load s3 config")
		}

		return cfg
	}).(S3Config)
}
