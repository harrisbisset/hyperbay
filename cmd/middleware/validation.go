package middleware

import (
	"net/http"

	"github.com/harrisbisset/webrelay/config"
	"github.com/harrisbisset/webrelay/toml"
)

func ValidateRelayConfig(handler config.RelayConfigHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := toml.ValidateRelayConfig(handler.GetRelayConfig()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		handler.ServeHTTP(w, r)
	}
}
