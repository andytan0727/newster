package main

import (
	"log"
	"net/http"

	"github.com/andytan0727/newster/api/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.StrictSlash(true)

	routes.SetupRoutes(r)

	log.Fatal(http.ListenAndServe(":8000", r))

}
