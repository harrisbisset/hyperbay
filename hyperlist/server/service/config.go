package service

import (
	"log"

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

	// optimise for future queries
	_, err := c.Write.Exec("PRAGMA optimize")
	if err != nil {
		log.Print(err)
	}
	_, err = c.Read.Exec("PRAGMA optimize")
	if err != nil {
		log.Print(err)
	}

	// close db connections
	err = c.Write.Close()
	if err != nil {
		log.Print(err)
	}
	err = c.Read.Close()
	if err != nil {
		log.Print(err)
	}
}

func NewConfig() Config {
	listHandler := toml.NewListHander()

	return Config{
		DBConfig:    database.CreateDBConfig(),
		ListHandler: listHandler,
	}
}
