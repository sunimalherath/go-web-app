package main

import (
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"

// Home - this is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome to the home page")
	if err != nil {
		log.Fatalln("Cannot load the home page", err)
	}
}

// About - this is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "This is the about page")
	if err != nil {
		log.Fatalln("Cannot load the about page", err)
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	http.ListenAndServe(portNumber, nil)
}
