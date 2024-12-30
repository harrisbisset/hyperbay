package middleware

import (
	"net/http"

	"github.com/harrisbisset/webrelay/config"
)

// checks if the hash header matches the expected value
// if no hash, then lets anything through
func AuthHash(handler config.RelayConfigHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if w.Header().Get("hash") != handler.GetRelayConfig().Hash {
			http.Error(w, "bad hash", http.StatusUnauthorized)
			return
		}
		handler.ServeHTTP(w, r)
	}
}
