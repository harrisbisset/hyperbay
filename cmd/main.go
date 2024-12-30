package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/harrisbisset/webrelay/config"
	"github.com/harrisbisset/webrelay/middleware"
	"github.com/harrisbisset/webrelay/routes"
	"github.com/harrisbisset/webrelay/toml"
)

func main() {
	cfg := config.NewConfig()
	mux := http.NewServeMux()

	// pages
	mux.Handle("GET /", routes.IndexHandler{RelayConfig: cfg.RelayConfig})

	// site routes
	site_len := len(cfg.Sites)
	var prev_site toml.Site
	var next_site toml.Site
	for i, s := range cfg.Sites {
		switch i {
		case 0:
			prev_site = cfg.Sites[site_len-1]
			next_site = cfg.Sites[i+1]
		case site_len - 1:
			prev_site = cfg.Sites[i-1]
			next_site = cfg.Sites[0]
		default:
			next_site = cfg.Sites[i+1]
			prev_site = cfg.Sites[i-1]
		}

		mux.Handle(fmt.Sprintf("GET /route/%s/next", s.Slug), routes.RouteHandler{Site: next_site})
		mux.Handle(fmt.Sprintf("GET /route/%s/prev", s.Slug), routes.RouteHandler{Site: prev_site})
	}

	// api
	mux.HandleFunc("GET /api/list", middleware.ValidateRelayConfig(routes.ListHandler{RelayConfig: cfg.RelayConfig}))
	mux.HandleFunc("GET /api/refresh", middleware.AuthHash(routes.RefreshHandler{RelayConfig: cfg.RelayConfig}))

	// OPTIONAL
	// http.HandleFunc("GET /random", middleware.AuthHash(routes.RandomHandler{RelayConfig: cfg.RelayConfig}))

	log.Fatal(http.ListenAndServe(":8080", middleware.NewLogger(mux, cfg.Logger)))
}
