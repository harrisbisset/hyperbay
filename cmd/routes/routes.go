package routes

import (
	"net/http"

	"github.com/harrisbisset/webrelay/toml"
)

type (
	// should respond to client with list of all sites, in json
	ListHandler struct {
		*toml.RelayConfig
	}

	// used to
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

}

func (handler RefreshHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func (handler RandomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
