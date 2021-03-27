package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/sunimalherath/go-web-app/pkg/config"
	"github.com/sunimalherath/go-web-app/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	//mux := pat.New()
	//
	//mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	//mux.Get("/about",http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()

	mux.Get("/",handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
