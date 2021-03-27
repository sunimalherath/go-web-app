package main

import (
	"github.com/bmizerany/pat"
	"github.com/sunimalherath/go-web-app/pkg/config"
	"github.com/sunimalherath/go-web-app/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about",http.HandlerFunc(handlers.Repo.About))

	return mux
}
