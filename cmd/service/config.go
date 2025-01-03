package service

import (
	"log"

	"github.com/harrisbisset/webrelay/service/toml"
)

type (
	Config struct {
		*toml.RelayConfig
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
		RelayConfig: relay,
		// Cache:       NewCacheSite(relay),
	}
}

// func NewCacheSite(relay *toml.RelayConfig) []SiteCache {
// 	now := time.Now().Unix()
// 	now := time.Now().Unix()
// 	cs := make([]SiteCache, len(relay.Sites))

// 	for i, v := range relay.Sites {
// 		cs[i] = SiteCache{
// 			Slug:    v.Slug,
// 			Alive:   true,
// 			Created: now,
// 			Expiry: ,
// 		}
// 	}
// 	return cs
// }
