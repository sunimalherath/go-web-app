package main

import (
	"fmt"
	"github.com/sunimalherath/go-web-app/pkg/config"
	"github.com/sunimalherath/go-web-app/pkg/handlers"
	"github.com/sunimalherath/go-web-app/pkg/render"
	"log"
	"net/http"
)

var app config.AppConfig

const portNumber = ":8080"

func main() {
	app.InProduction = false
	tc, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Println("Unable to create template cache, ", err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)

	//http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
