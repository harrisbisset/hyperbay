package toml

import (
	"errors"
	"os"

	"github.com/pelletier/go-toml"
)

type (
	Site struct {
		Id      int
		Slug    string `json:"slug" toml:"slug"`
		Name    string `json:"name" toml:"name"`
		Src     string `json:"src" toml:"src"`
		Url     string `json:"url" toml:"url"`
		Created int    `json:"created" toml:"created"`
		Alive   bool   `json:"alive" toml:"alive"`
	}
)

const TOML_ERROR_NOT_FOUND = "relay.toml not found"

func ParseHyperlist() (listConfig, error) {
	var cfg listConfig

	doc, err := os.ReadFile("./relay.toml")
	if err != nil {
		return cfg, errors.New(TOML_ERROR_NOT_FOUND)
	}

	err = toml.Unmarshal([]byte(doc), &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
