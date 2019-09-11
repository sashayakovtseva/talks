package main

import (
	"flag"
	"log"
	"net"
)

var addr string

func init() {
	flag.StringVar(&addr, "addr", ":8080", "server address to listen")
}

func main() {
	flag.Parse()

	log.Printf("Listening on %s", addr)
	// ln, err := net.Listen("tcp6", addr)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Could not listen: %v", err)
	}
	ln.Close()
}
