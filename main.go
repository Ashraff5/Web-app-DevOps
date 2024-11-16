package main

import (
	"log"
	"net/http"
)

func about(w http.ResponseWriter, r *http.Request) {
	// Render the home html page from static folder
	http.ServeFile(w, r, "static/about.html")
}

func project(w http.ResponseWriter, r *http.Request) {
	// Render the course html page
	http.ServeFile(w, r, "static/project.html")
}


func main() {

	http.HandleFunc("/about", about)
	http.HandleFunc("/project", project)
	

	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
