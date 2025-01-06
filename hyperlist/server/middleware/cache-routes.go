package middleware

// type SiteMemoryCacher struct {
// 	lock  sync.RWMutex
// 	items map[string]*struct {
// 		Slug    string
// 		Alive   bool
// 		Created int64
// 		Expiry  int64
// 	}
// 	interval int // GC interval
// }

// func CacheSiteRoute(cache config.SiteCache) func(handler http.Handler) http.Handler {
// 	return func(handler http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			cache.Expiry > 0 &&
// 				(time.Now().Unix()-item.created) >= item.expire

// 		})
// 	}
// }
