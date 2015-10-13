package main

import (
	"log"
	"net/http"
	"path/filepath"
	//"fmt"
)

func main() {
	directory := "./templates"
	// Simple static webserver:

	static_path := filepath.Join(directory, "static/")
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(static_path))))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Render(directory, w, r)
	})

	log.Fatal(http.ListenAndServe(":8080", mux))

}
