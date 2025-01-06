package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/harrisbisset/hyperbay/hyperlist/server/service/toml"
)

type (
	// should respond to client with list of all sites, in json
	ListHandler struct {
		*toml.ListHandler
	}

	// used to refresh the site list
	// does not refresh any other data
	RefreshHandler struct {
		*toml.ListHandler
	}

	// used to get a "random" site from the server
	RandomHandler struct {
		*toml.ListHandler
	}
)

func (handler ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(handler.Sites())
	if err != nil {
		log.Print(err)
		http.Error(w, "failed to marshal json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (handler RefreshHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := handler.ListHandler.RefreshHyperlist(); err != nil {
		log.Print(err)
		http.Error(w, "parse failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler RandomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	site, err := handler.RandomSite()
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, site.Src, http.StatusPermanentRedirect)
}
