package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var addr string

func init() {
	flag.StringVar(&addr, "addr", ":8080", "server address to listen")
}

func main() {
	flag.Parse()

	log.Printf("Listening on %s", addr)
	err := http.ListenAndServe(addr, http.HandlerFunc(hello))
	log.Fatalf("Server error: %v", err)
}

func hello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello, Go!")
}
