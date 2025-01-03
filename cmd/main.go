package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/harrisbisset/webrelay/middleware"
	"github.com/harrisbisset/webrelay/routes"
	"github.com/harrisbisset/webrelay/routes/api"
	"github.com/harrisbisset/webrelay/service"
)

func main() {
	cfg := service.NewConfig()
	r := chi.NewRouter()

	// middlewares
	r.Use(chimiddleware.Logger)

	// pages
	r.Get("/", routes.IndexHandler{RelayConfig: cfg.RelayConfig}.ServeHTTP)
	r.Get("/list", routes.ListHandler{RelayConfig: cfg.RelayConfig}.ServeHTTP)

	for _, v := range service.GetRing(cfg.Sites) {
		r.Get(v.Path, routes.RouteHandler{Site: v.Site}.ServeHTTP)
	}

	// api
	r.Group(func(r chi.Router) {
		r.Get("/api/list", api.ListHandler{RelayConfig: cfg.RelayConfig}.ServeHTTP)
		r.Get("/api/random", api.RandomHandler{RelayConfig: cfg.RelayConfig}.ServeHTTP)

		// auth
		r.Group(func(r chi.Router) {
			r.Use(middleware.AuthHash(cfg.Hash))
			r.Get("/api/refresh", api.RefreshHandler{RelayConfig: cfg.RelayConfig}.ServeHTTP)
		})
	})

	http.ListenAndServe(":8080", r)
}
