package config

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/harrisbisset/webrelay/toml"
)

type (
	Config struct {
		*toml.RelayConfig
		*slog.Logger
	}

	RelayConfigHandler interface {
		http.Handler
		GetRelayConfig() *toml.RelayConfig
	}
)

func NewConfig() Config {
	relay, err := toml.ParseRelay()
	if err != nil {
		panic(err)
	}

	if relay.Hash == "" {
		log.Print("WARNING: hash is empty")
	}

	return Config{
		Logger:      slog.Default(),
		RelayConfig: relay,
	}
}
