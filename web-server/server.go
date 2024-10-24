package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	mux := http.NewServeMux()
	serveStatic(mux)
	// serveNgWithoutRoute(mux)
	serveNgWithRoute(mux)

	port := "8080"
	log.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))

}

func serveStatic(r *http.ServeMux) {
	staticDir := "./static"
	fs := http.FileServer(http.Dir(staticDir))
	r.Handle("/web/", http.StripPrefix("/web/", fs))
}

func serveNgWithoutRoute(r *http.ServeMux) {
	distDir := "./ng-without-route"
	fs := http.FileServer(http.Dir(distDir))
	r.Handle("/", fs)
}

func serveNgWithRoute(r *http.ServeMux) {
	distDir := "./ng-with-route"
	indexFile := filepath.Join(distDir, "index.html")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestedPath := filepath.Join(distDir, r.URL.Path)
		log.Printf("URL Path: %s", r.URL.Path)
		if _, err := os.Stat(requestedPath); os.IsNotExist(err) {
			log.Printf("RequestedPath not exist: %s", requestedPath)
			http.ServeFile(w, r, indexFile)
		} else {
			http.ServeFile(w, r, requestedPath)
		}
	})
}
