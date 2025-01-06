package toml

import (
	"sync"
	"time"
)

type (
	SiteCache struct {
		mu      sync.Mutex
		handler *ListHandler

		// time since last refresh
		lastRefresh int64

		// when site should automatically refresh
		refreshTime int64

		sites []SiteInfo
	}

	SiteInfo struct {
		Id      int
		Created int64
		Alive   bool
	}
)

func NewCacheSite(handler *ListHandler, sites []Site) *SiteCache {
	now := time.Now().Unix()

	c := &SiteCache{
		handler:     handler,
		lastRefresh: now,
		sites:       make([]SiteInfo, len(sites)),
	}

	for _, site := range c.sites {

		// created should never be zero, unless it was never set in the toml
		// so we should give it a value
		var created int64
		if site.Created == 0 {
			created = now
		}

		site = SiteInfo{
			Id:      site.Id,
			Alive:   site.Alive,
			Created: created,
		}
	}

	return c
}
