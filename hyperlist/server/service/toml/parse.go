package toml

import (
	"os"

	"github.com/pelletier/go-toml"
)

type (
	Site struct {
		Slug string `json:"slug" toml:"slug"`
		Name string `json:"name" toml:"name"`
		Src  string `json:"src" toml:"src"`
		Url  string `json:"url" toml:"url"`
	}

	RelayConfig struct {
		Hash      string `toml:"hash"`
		HostUser  string `toml:"hostUser"`
		HostEmail string `toml:"hostEmail"`
		Sites     []Site `toml:"sites"`
	}
)

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
