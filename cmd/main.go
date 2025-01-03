package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/harrisbisset/webrelay/config"
	"github.com/harrisbisset/webrelay/middleware"
	"github.com/harrisbisset/webrelay/routes"
	"github.com/harrisbisset/webrelay/service"
)

func main() {
	cfg := config.NewConfig()
	r := chi.NewRouter()

	// middlewares
	r.Use(chimiddleware.Logger)

	// pages
	r.Get("/", routes.IndexHandler{RelayConfig: cfg.RelayConfig}.ServeHTTP)

	for _, v := range service.GetRing(cfg.Sites) {
		r.Get(v.Path, routes.RouteHandler{Site: v.Site}.ServeHTTP)
	}

	// api
	r.Group(func(r chi.Router) {
		r.Get("/api/list", routes.ListHandler{RelayConfig: cfg.RelayConfig}.ServeHTTP)
		r.Get("/api/random", routes.RandomHandler{RelayConfig: cfg.RelayConfig}.ServeHTTP)

		// auth
		r.Group(func(r chi.Router) {
			r.Use(middleware.AuthHash(cfg.Hash))
			r.Get("/api/refresh", routes.RefreshHandler{RelayConfig: cfg.RelayConfig}.ServeHTTP)
		})
	})

	http.ListenAndServe(":8080", r)
}
