package routes

import (
	"net/http"
	"text/template"

	"github.com/harrisbisset/webrelay/toml"
)

type (
	IndexHandler struct {
		*toml.RelayConfig
	}

	RouteHandler struct {
		toml.Site
	}
)

func (handler IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./cmd/templates/index.go.html"))
	tmpl.Execute(w, *handler.RelayConfig)
}

func (handler RouteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, handler.Src, http.StatusPermanentRedirect)
}
