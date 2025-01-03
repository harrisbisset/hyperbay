package routes

import (
	"net/http"
	"text/template"

	"github.com/harrisbisset/webrelay/service/toml"
)

type (
	IndexHandler struct {
		*toml.RelayConfig
	}

	ListHandler struct {
		*toml.RelayConfig
	}

	RouteHandler struct {
		toml.Site
	}
)

func (handler IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/index.go.html"))
	tmpl.Execute(w, *handler.RelayConfig)
}

func (handler ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/index.go.html"))
	tmpl.Execute(w, *handler.RelayConfig)
}

func (handler RouteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, handler.Site.Src, http.StatusPermanentRedirect)
}
