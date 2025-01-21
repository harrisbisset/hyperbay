package service

import (
	"github.com/harrisbisset/hyperbay/hyperlist/server/service/cache"
	"github.com/harrisbisset/hyperbay/hyperlist/server/service/database"
	"github.com/harrisbisset/hyperbay/hyperlist/server/service/toml"
)

type (
	Config struct {
		*cache.Cache
		database.DBConfig
	}
)

func (cfg Config) Close() {
	cfg.DBConfig.Close()
}

func NewConfig() Config {
	listHandler := toml.NewCache()

	return Config{
		DBConfig:    database.CreateDBConfig(),
		ListHandler: listHandler,
	}
}
