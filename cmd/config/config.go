package config

import "github.com/harrisbisset/webrelay/toml"

type Config struct {
	*toml.RelayConfig
}

func NewConfig() Config {
	relay, err := toml.ParseRelay()
	if err != nil {
		panic(err)
	}

	return Config{
		RelayConfig: relay,
	}
}
