package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/sunimalherath/go-web-app/pkg/config"
	"github.com/sunimalherath/go-web-app/pkg/handlers"
	"github.com/sunimalherath/go-web-app/pkg/render"
	"log"
	"net/http"
	"time"
)

var app config.AppConfig
var session *scs.SessionManager

const portNumber = ":8080"

func main() {
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

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
