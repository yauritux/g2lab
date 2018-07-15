package main

import (
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var port = flag.String("port", ":8000", "HTTP service port")

func main() {
	flag.Parse()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(*port, router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
