package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello, Friend!")
		if err != nil {
			log.Fatalln(err)
		}
	})

	http.ListenAndServe(":8080", nil)
}
