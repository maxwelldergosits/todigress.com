package main

import (
	"log"
	"net/http"
	"path/filepath"
  "os"
  "todigress-server/templates"
)

func main() {
  var directory string

  if len(os.Args) < 2 {
	  directory = GetDefaultConf().Directory()
  } else {
	  directory = os.Args[1]
  }

	// Simple static webserver:
	static_path := filepath.Join(directory, "static/")
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(static_path))))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates.Render(directory, w, r)
	})

	log.Fatal(http.ListenAndServe(":8080", mux))

}
