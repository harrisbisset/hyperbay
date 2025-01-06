package service

import (
	"github.com/harrisbisset/hyperbay/hyperlist/server/service/toml"
)

type (
	Config struct {
		*toml.ListHandler
	}
)

func NewConfig() Config {
	return Config{
		ListHandler: toml.NewListHander(),
	}
}
