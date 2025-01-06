package toml

import (
	"errors"
	"iter"
	"math/rand/v2"
	"sync"
)

type (
	Site struct {
		Id      int    `toml:"id"`
		Slug    string `toml:"slug"`
		Name    string `toml:"name"`
		Src     string `toml:"src"`
		Url     string `toml:"url"`
		Created int64
		Alive   bool
	}

	tomlConfig struct {
		mu sync.Mutex

		hash      string `toml:"hash"`
		hostUser  string `toml:"hostUser"`
		hostEmail string `toml:"hostEmail"`
		sites     []Site
	}

	ListHandler struct {
		*tomlConfig
		cache *SiteCache
	}
)

func (cfg *tomlConfig) Hash() string {
	return cfg.hash
}

func (cfg *tomlConfig) HostUser() string {
	return cfg.hostUser
}

func (cfg *tomlConfig) HostEmail() string {
	return cfg.hostEmail
}

// gets a random alive site from the site list
// errors if no alive sites are found
func (cfg *tomlConfig) RandomSite() (Site, error) {
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

func (handler *ListHandler) IterSites() iter.Seq[Site] {
	return func(yield func(Site) bool) {
		for _, site := range handler.sites {
			if !yield(site) {
				return
			}
		}
	}
}

func (handler *ListHandler) RefreshHyperlist() error {
	handler.tomlConfig.mu.Lock()
	handler.cache.mu.Lock()
	defer handler.tomlConfig.mu.Unlock()
	defer handler.cache.mu.Unlock()

	// todo

	return nil
}

func NewListHander() *ListHandler {
	cfg, err := ParseHyperlist()
	if err != nil {
		panic(err)
	}

	return &ListHandler{
		tomlConfig: cfg,
		cache:      NewCacheSite(cfg.sites),
	}
}
