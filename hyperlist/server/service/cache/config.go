package cache

import (
	"errors"
	"math/rand/v2"
)

// gets a random alive site from the site list
// errors if no alive sites are found
func (cfg *Cache) RandomSite() (Site, error) {
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

// func (handler *ListHandler) RefreshHyperlist() error {
// 	handler.tomlConfig.mu.Lock()
// 	handler.cache.mu.Lock()
// 	defer handler.tomlConfig.mu.Unlock()
// 	defer handler.cache.mu.Unlock()

// 	// todo

// 	return nil
// }

// func NewListHander() *ListHandler {
// 	cfg, err := ParseHyperlist()
// 	if err != nil {
// 		panic(err)
// 	}

// 	return &ListHandler{
// 		tomlConfig: cfg,
// 		cache:      NewCacheSite(cfg.sites),
// 	}
// }
