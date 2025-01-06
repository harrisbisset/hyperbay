package routes

import (
	"net/http"

	"github.com/a-h/templ"
)

func Render(c templ.Component, w http.ResponseWriter, r *http.Request) {
	templ.Handler(c).ServeHTTP(w, r)
}
