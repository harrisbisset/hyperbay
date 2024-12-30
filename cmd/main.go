package main

import (
	"net/http"

	"github.com/harrisbisset/webrelay/config"
	"github.com/harrisbisset/webrelay/routes"
)

func main() {
	cfg := config.NewConfig()

	http.Handle("GET /list", routes.ListHandler{RelayConfig: cfg.RelayConfig})
	http.Handle("GET /refresh", routes.RefreshHandler{RelayConfig: cfg.RelayConfig})

	// OPTIONAL
	// http.Handle("GET /random", routes.RandomHandler{})

	http.ListenAndServe(":8080", nil)
}
