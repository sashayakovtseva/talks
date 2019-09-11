package main

import (
	"flag"
	"log"
	"net"
)

var (
	addr string
	netw string
)

func init() {
	flag.StringVar(&addr, "addr", ":8080", "server address to listen")
	flag.StringVar(&netw, "net", "tcp", "network to use")
}

func main() {
	flag.Parse()

	log.Printf("Listening on %s", addr)
	ln, err := net.Listen(netw, addr)
	if err != nil {
		log.Fatalf("Could not listen: %v", err)
	}
	ln.Close()
}
