package toml

import (
	"errors"
	"os"

	"github.com/pelletier/go-toml"
)

type (
	TomlConfig struct {
		Hash      string `toml:"hash"`
		HostUser  string `toml:"hostUser"`
		HostEmail string `toml:"hostEmail"`
		Sites     []struct {
			Slug string `toml:"slug"`
			Name string `toml:"name"`
			Src  string `toml:"src"`
			Url  string `toml:"url"`
		} `toml:"sites"`
	}
)

func ParseHyperlist() (TomlConfig, error) {
	var cfg TomlConfig

	doc, err := os.ReadFile("./hyperlist.toml")
	if err != nil {
		return cfg, errors.New("hyperlist.toml not found")
	}
	if err = toml.Unmarshal(doc, &cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}
