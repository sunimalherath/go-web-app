package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// Home - this is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

// About - this is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	// 1. Parse file
	parsedTmpl, _ := template.ParseFiles("./templates/" + tmpl)

	// 2. Execute parsed file
	err := parsedTmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
	}
}
