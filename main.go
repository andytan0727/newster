package main

import (
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/andytan0727/newster/api/routes"
	"github.com/gorilla/mux"
)

type spaHandler struct {
	staticPath string
	indexPath  string
}

// ServeHTTP serves the index.html from frontend SPA
// dist/ directory.
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// routing and 404 Error page is handled on the client side
	http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
}

func main() {
	r := mux.NewRouter()
	r.StrictSlash(true)

	// setup api routes
	routes.SetupRoutes(r)

	// setup SPA frontend route
	spa := spaHandler{staticPath: "frontend/dist", indexPath: "index.html"}
	r.PathPrefix("/").Handler(spa)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
