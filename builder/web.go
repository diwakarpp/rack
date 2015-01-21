package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/convox/kernel/builder/Godeps/_workspace/src/github.com/codegangsta/negroni"
	"github.com/convox/kernel/builder/Godeps/_workspace/src/github.com/gorilla/mux"
	"github.com/convox/kernel/builder/controllers/builder"
)

var port string = "3000"

func redirect(path string) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		http.Redirect(rw, r, "/clusters", http.StatusFound)
	}
}

func parseForm(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	r.ParseMultipartForm(2048)
	next(rw, r)
}

func main() {
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	router := mux.NewRouter()

	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello, World!"))
	}).Methods("GET")

	router.HandleFunc("/clusters/{cluster}/apps/{app}/build", builder.Build).Methods("POST")

	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(parseForm))
	n.UseHandler(router)
	n.Run(fmt.Sprintf(":%s", port))
}
