package service

import (
	"github.com/harrisbisset/hyperbay/hyperlist/server/service/database"
	"github.com/harrisbisset/hyperbay/hyperlist/server/service/toml"
)

type (
	Config struct {
		*toml.ListHandler
		database.DBConfig
	}
)

func (c Config) Close() {
	c.DB.Close()
}

func NewConfig() Config {
	listHandler := toml.NewListHander()

	return Config{
		DBConfig:    database.CreateDBConfig(),
		ListHandler: listHandler,
	}
}
