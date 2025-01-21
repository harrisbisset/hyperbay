package routes

import (
	"net/http"

	"github.com/harrisbisset/hyperbay/hyperlist/server/render/views/view_index"
	"github.com/harrisbisset/hyperbay/hyperlist/server/render/views/view_list"
	"github.com/harrisbisset/hyperbay/hyperlist/server/service/cache"
)

type (
	IndexHandler struct {
		*cache.Cache
	}

	ListHandler struct {
		*cache.Cache
	}
)

func (handler IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	Render(view_index.Show(*handler.Cache), w, r)
}

func (handler ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	Render(view_list.Show(*handler.Cache), w, r)
}
