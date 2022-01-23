package main

import (
	"log"
	"net/http"

	"github.com/algren123/go-library/packages/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterLibraryRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
