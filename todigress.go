package main

import (
	"log"
	"net/http"
	"path/filepath"
  "os"
  _ "net/http/pprof"
  "bytes"
  "io"
  "time"
)
type CacheEntry struct {
  responsePayload []byte
  expiresAt time.Time
}

func main() {
  var directory string

  if len(os.Args) < 2 {
	  directory = GetDefaultConf().Directory()
  } else {
	  directory = os.Args[1]
  }

  go func() {
  log.Println(http.ListenAndServe(":6060", nil))
  }()

	// Simple static webserver:
	static_path := filepath.Join(directory, "static/")
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(static_path))))

  cache := make(map[string]CacheEntry)

	mux.HandleFunc("/", func(httpW http.ResponseWriter, r *http.Request) {

    if val, ok := cache[r.RequestURI]; ok {
      if (time.Now().Before(val.expiresAt)) {
        httpW.Write(val.responsePayload)
        return
      } else {
        delete(cache, r.RequestURI)
      }
    }

    var cacheBuffer bytes.Buffer

    writer := io.MultiWriter(httpW, &cacheBuffer);
		Render(directory, writer, r)

    cache[r.RequestURI]= CacheEntry{cacheBuffer.Bytes(), time.Now().Add(3 * time.Second)}

	})

	log.Fatal(http.ListenAndServe(":8080", mux))

}
