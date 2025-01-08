package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/harrisbisset/hyperbay/hyperlist/server/middleware"
	"github.com/harrisbisset/hyperbay/hyperlist/server/routes"
	"github.com/harrisbisset/hyperbay/hyperlist/server/routes/api"
	"github.com/harrisbisset/hyperbay/hyperlist/server/service"
)

func main() {
	cfg := service.NewConfig()
	defer cfg.Close()
	r := chi.NewRouter()

	// middlewares
	r.Use(chimiddleware.Logger)

	// pages
	r.Get("/", routes.IndexHandler{ListHandler: cfg.ListHandler}.ServeHTTP)
	r.Get("/list", routes.ListHandler{ListHandler: cfg.ListHandler}.ServeHTTP)

	// api
	r.Group(func(r chi.Router) {
		r.Get("/api/list", api.ListHandler{ListHandler: cfg.ListHandler}.ServeHTTP)
		r.Get("/api/random", api.RandomHandler{ListHandler: cfg.ListHandler}.ServeHTTP)

		// auth
		r.Group(func(r chi.Router) {
			r.Use(middleware.AuthHash(cfg.Hash()))
			r.Get("/api/refresh", api.RefreshHandler{ListHandler: cfg.ListHandler}.ServeHTTP)
		})
	})

	r.Handle("/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir("./render/public"))))

	http.ListenAndServe(":80", r)
}
