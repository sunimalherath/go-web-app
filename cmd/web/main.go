package main

import (
	"fmt"
	"net/http"

	"github.com/sunimalherath/go-web-app/pkg/config"
	"github.com/sunimalherath/go-web-app/pkg/handlers"
	"github.com/sunimalherath/go-web-app/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Println("Unable to create template cache, ", err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	http.ListenAndServe(portNumber, nil)
}
