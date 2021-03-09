package handlers

import (
	"net/http"

	"github.com/sunimalherath/go-web-app/pkg/render"
)

// Home - this is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About - this is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
