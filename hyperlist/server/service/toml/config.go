package toml

type (
	listConfig struct {
		hash      string `toml:"hash"`
		hostUser  string `toml:"hostUser"`
		hostEmail string `toml:"hostEmail"`
		sites     []Site `toml:"sites"`
	}

	ListHandler struct {
		listConfig
		cache *SiteCache
	}
)

func NewListHander() *ListHandler {
	handler := &ListHandler{}

	cfg, err := ParseHyperlist()
	if err != nil {
		panic(err)
	}

	handler.listConfig = cfg
	handler.cache = NewCacheSite(handler, cfg.sites)

	return handler
}
