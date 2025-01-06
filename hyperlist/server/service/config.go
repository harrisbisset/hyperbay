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
	listHandler := toml.NewListHander()

	return Config{
		ListHandler: listHandler,
	}
}
