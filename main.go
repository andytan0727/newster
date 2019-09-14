package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/andytan0727/newster/api/routes"
	"github.com/gorilla/mux"
)

type spaHandler struct {
	staticPath string
	indexPath  string
}

// ServeHTTP serves static files from frontend SPA
// dist/ directory.
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(h.staticPath, r.URL.Path)

	// check whether a file exists at the given path
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// serve static dir otherwise
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func main() {
	port := os.Getenv("PORT") // to get port defined by hosting platform

	if port == "" {
		port = "8000"
	}

	r := mux.NewRouter()
	r.StrictSlash(true)

	// setup api routes
	routes.SetupRoutes(r)

	// setup SPA frontend route
	spa := spaHandler{staticPath: "frontend/dist", indexPath: "index.html"}
	r.PathPrefix("/").Handler(spa)

	srv := &http.Server{
		Handler: r,

		// using 0.0.0.0 to allow access from external world
		// especially with the usage of docker
		Addr: "0.0.0.0:" + port,

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
