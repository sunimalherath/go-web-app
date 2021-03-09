package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func renderTemplate(w http.ResponseWriter, tmpl string) {
	// 1. Parse file
	parsedTmpl, _ := template.ParseFiles("./templates/" + tmpl)

	// 2. Execute parsed file
	err := parsedTmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
	}
}
