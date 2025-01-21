package cache

import (
	"sync"
	"time"

	"github.com/harrisbisset/hyperbay/hyperlist/server/service/toml"
)

type (
	Cache struct {
		mu sync.Mutex

		// time since last refresh
		lastRefresh int64

		// when site should automatically refresh
		refreshTime int64

		hash      string
		hostUser  string
		hostEmail string
		sites     []Site
	}

	Site struct {
		Id      int
		Slug    string
		Name    string
		Src     string
		Url     string
		Created int64
		Alive   bool
	}
)

func (c *Cache) UpdateCache() {

}

func NewCache() (*Cache, error) {
	cfg, err := toml.ParseHyperlist()
	if err != nil {
		return nil, err
	}

	return parseToml(cfg), nil
}

func NewCacheSite(sites []Site) *Cache {
	now := time.Now().Unix()

	c := &Cache{
		lastRefresh: now,
		sites:       make([]Site, len(sites)),
	}

	for _, site := range c.sites {

		// created should never be zero, unless it was never set in the toml
		// so we should give it a value
		var created int64
		if site.Created == 0 {
			created = now
		}

		site = Site{
			Id:      site.Id,
			Alive:   site.Alive,
			Created: created,
		}
	}

	return c
}

func parseToml(cfg toml.TomlConfig) *Cache {
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

	return &Cache{
		hash:      cfg.Hash,
		hostUser:  cfg.HostUser,
		hostEmail: cfg.HostEmail,
		sites:     sites,
	}
}
