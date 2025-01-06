package toml

import (
	"errors"
	"math/rand/v2"
	"sync"
)

type (
	listConfig struct {
		mu sync.Mutex

		hash      string `toml:"hash"`
		hostUser  string `toml:"hostUser"`
		hostEmail string `toml:"hostEmail"`
		sites     []Site `toml:"sites"`
	}

	ListHandler struct {
		*listConfig
		cache *SiteCache
	}
)

func (cfg *listConfig) Hash() string {
	return cfg.hash
}

func (cfg *listConfig) HostUser() string {
	return cfg.hostUser
}

func (cfg *listConfig) HostEmail() string {
	return cfg.hostEmail
}

// gets a random alive site from the site list
// errors if no alive sites are found
func (cfg *listConfig) RandomSite() (Site, error) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	l := len(cfg.sites) - 1
	if l <= 0 {
		return Site{}, errors.New("cannot get random site, not enough sites")
	}
	rn := rand.IntN(l)

	count := 0
	var s Site
	for {
		// if full rotation has been completed, and no sites are available
		if count == l {
			return Site{}, errors.New("no alive sites found")
		}

		s = cfg.sites[rn]

		// if site link isn't dead
		if s.Alive {
			return s, nil
		}

		// move onto next site in list
		rn++
		count++
		if rn > l {
			rn = 0
		}
	}
}

func (handler *ListHandler) Sites() []Site {
	return handler.sites
}

func (handler *ListHandler) RefreshHyperlist() error {
	handler.listConfig.mu.Lock()
	handler.cache.mu.Lock()
	defer handler.listConfig.mu.Unlock()
	defer handler.cache.mu.Unlock()

	return nil
}

func NewListHander() *ListHandler {
	cfg, err := ParseHyperlist()
	if err != nil {
		panic(err)
	}

	return &ListHandler{
		listConfig: cfg,
		cache:      NewCacheSite(cfg.sites),
	}
}
