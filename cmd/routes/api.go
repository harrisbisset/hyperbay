package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/harrisbisset/webrelay/toml"
)

type (
	// should respond to client with list of all sites, in json
	ListHandler struct {
		*toml.RelayConfig
	}

	// used to refresh the site list
	// does not refresh any other data
	RefreshHandler struct {
		*toml.RelayConfig
	}

	// NOT RECOMMENDED
	// used to get a random site from the server
	RandomHandler struct {
		*toml.RelayConfig
	}
)

func (handler ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(handler.Sites)
	if err != nil {
		log.Print(err)
		http.Error(w, "failed to marshal json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (handler ListHandler) GetRelayConfig() *toml.RelayConfig {
	return handler.RelayConfig
}

func (handler RefreshHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	relay, err := toml.ParseRelay()
	if err != nil {
		log.Print(err)
		http.Error(w, "parse failed", http.StatusInternalServerError)
		return
	}

	// update relay and respond
	handler.RelayConfig.Sites = relay.Sites
	w.WriteHeader(http.StatusOK)
}

func (handler RefreshHandler) GetRelayConfig() *toml.RelayConfig {
	return handler.RelayConfig
}

func (handler RandomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func (handler RandomHandler) GetRelayConfig() *toml.RelayConfig {
	return handler.RelayConfig
}
