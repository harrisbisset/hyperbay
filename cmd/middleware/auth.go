package middleware

import (
	"net/http"
)

// checks if the hash header matches the expected value
// if no hash, then lets anything through
func AuthHash(hash string) func(handler http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if w.Header().Get("hash") != hash {
				http.Error(w, "bad hash", http.StatusUnauthorized)
				return
			}
		})
	}
}
