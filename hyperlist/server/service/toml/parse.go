package toml

import (
	"errors"
	"os"

	"github.com/pelletier/go-toml"
)

type (
	tomlUserConfig struct {
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

const TOML_ERROR_NOT_FOUND = "hyperlist.toml not found"

func ParseHyperlist() (*tomlConfig, error) {
	doc, err := os.ReadFile("./hyperlist.toml")
	if err != nil {
		return nil, errors.New("hyperlist.toml not found")
	}

	var cfg tomlUserConfig
	err = toml.Unmarshal(doc, &cfg)
	if err != nil {
		return nil, err
	}

	var sites []Site
	for i, s := range cfg.Sites {
		sites = append(sites, Site{
			Id:   i,
			Slug: s.Slug,
			Name: s.Name,
			Src:  s.Src,
			Url:  s.Url,
		})
	}

	return &tomlConfig{
		hash:      cfg.Hash,
		hostUser:  cfg.HostUser,
		hostEmail: cfg.HostEmail,
		sites:     sites,
	}, nil
}
