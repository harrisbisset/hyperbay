package toml

import (
	"os"

	"github.com/pelletier/go-toml"
)

type RelayConfig struct {
	Sites []struct {
		Src  string `toml:"src"`
		Name string `toml:"name"`
	} `toml:"sites"`
}

func ParseRelay() (*RelayConfig, error) {
	doc, err := os.ReadFile("./relay.toml")
	if err != nil {
		panic("relay.toml not found")
	}

	var cfg RelayConfig
	err = toml.Unmarshal([]byte(doc), &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
