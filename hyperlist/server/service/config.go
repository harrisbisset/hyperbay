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

func (cfg Config) Close() {
	cfg.DBConfig.Close()
}

func NewConfig() Config {
	listHandler := toml.NewListHander()

	return Config{
		DBConfig:    database.CreateDBConfig(),
		ListHandler: listHandler,
	}
}
