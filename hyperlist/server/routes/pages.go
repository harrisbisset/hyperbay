package routes

import (
	"net/http"

	"github.com/harrisbisset/hyperbay/hyperlist/server/render/views/view_index"
	"github.com/harrisbisset/hyperbay/hyperlist/server/render/views/view_list"
	"github.com/harrisbisset/hyperbay/hyperlist/server/service/toml"
)

type (
	IndexHandler struct {
		*toml.ListHandler
	}

	ListHandler struct {
		*toml.ListHandler
	}
)

func (handler IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	Render(view_index.Show(*handler.ListHandler), w, r)
}

func (handler ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	Render(view_list.Show(*handler.ListHandler), w, r)
}
