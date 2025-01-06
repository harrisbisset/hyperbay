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

const TOML_ERROR_NOT_FOUND = "hyperlist.toml not found"

func ParseHyperlist() (*listConfig, error) {
	doc, err := os.ReadFile("./hyperlist.toml")
	if err != nil {
		return nil, errors.New(TOML_ERROR_NOT_FOUND)
	}

	var cfg listConfig
	err = toml.Unmarshal([]byte(doc), &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
