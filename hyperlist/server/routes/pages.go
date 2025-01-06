package routes

import (
	"net/http"

	"github.com/harrisbisset/hyperbay/hyperlist/server/render/views/view_index"
	"github.com/harrisbisset/hyperbay/hyperlist/server/service/toml"
)

type (
	IndexHandler struct {
		*toml.ListHandler
	}

	ListHandler struct {
		*toml.ListHandler
	}

	RouteHandler struct {
		toml.Site
	}
)

func (handler IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	Render(view_index.Show(), w, r)
}

func (handler ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	Render(view_index.Show(), w, r)
}

func (handler RouteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	Render(view_index.Show(), w, r)
}
