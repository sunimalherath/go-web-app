package handlers

import (
	"github.com/sunimalherath/go-web-app/pkg/config"
	"net/http"

	"github.com/sunimalherath/go-web-app/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository fo the handelers.
func NewHandlers(r *Repository) {
	Repo = r
}

// Home - this is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About - this is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
