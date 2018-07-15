package main

import (
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
)

var addr = flag.String("address", ":8000", "HTTP service port")

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(*addr, nil))
}
