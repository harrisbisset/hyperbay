package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"

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

	{ // routes

		//pages
		r.Get("/", routes.IndexHandler{Cache: cfg.Cache}.ServeHTTP)
		r.Get("/list", routes.ListHandler{Cache: cfg.Cache}.ServeHTTP)

		// api
		r.Group(func(r chi.Router) {
			r.Get("/api/list", api.ListHandler{Cache: cfg.Cache}.ServeHTTP)
			r.Get("/api/random", api.RandomHandler{Cache: cfg.Cache}.ServeHTTP)

			// auth
			r.Group(func(r chi.Router) {
				r.Use(middleware.AuthHash(cfg.Hash()))
				r.Get("/api/refresh", api.RefreshHandler{Cache: cfg.Cache}.ServeHTTP)
			})
		})

		// static
		r.Handle("/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir("./render/public"))))
	}

	server := http.Server{
		Addr:    ":80",
		Handler: r,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen and serve returned err: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("got interruption signal")
	if err := server.Shutdown(context.TODO()); err != nil {
		log.Printf("server shutdown returned an err: %v\n", err)
	}
}
