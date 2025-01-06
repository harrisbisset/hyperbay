package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/harrisbisset/hyperbay/hyperlist/server/service"
)

func main() {
	_ = service.NewConfig()
	r := chi.NewRouter()

	// middlewares
	r.Use(chimiddleware.Logger)

	// pages
	// r.Get("/", routes.IndexHandler{RelayConfig: cfg.RelayConfig}.ServeHTTP)
	// r.Get("/list", routes.ListHandler{RelayConfig: cfg.RelayConfig}.ServeHTTP)

	// for _, v := range service.GetRing(cfg.Sites) {
	// 	r.Get(v.Path, routes.RouteHandler{Site: v.Site}.ServeHTTP)
	// }

	// // api
	// r.Group(func(r chi.Router) {
	// 	r.Get("/api/list", api.ListHandler{RelayConfig: cfg.RelayConfig}.ServeHTTP)
	// 	r.Get("/api/random", api.RandomHandler{RelayConfig: cfg.RelayConfig}.ServeHTTP)

	// 	// auth
	// 	r.Group(func(r chi.Router) {
	// 		r.Use(middleware.AuthHash(cfg.Hash))
	// 		r.Get("/api/refresh", api.RefreshHandler{RelayConfig: cfg.RelayConfig}.ServeHTTP)
	// 	})
	// })

	http.ListenAndServe(":80", r)
}
