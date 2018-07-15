package main

import (
	"flag"
	"log"
	"net/http"
)

var port = flag.String("port", ":8000", "HTTP Service Port")

func main() {
	flag.Parse()
	router := NewRouter()

	log.Fatal(http.ListenAndServe(*port, router))
}
