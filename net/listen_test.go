package main

import (
	"net"
	"testing"
)

func BenchmarkListen(b *testing.B) {
	tt := []struct {
		netw string
		addr []string
	}{
		{
			netw: "tcp",
			addr: []string{"127.0.0.1:8080", ":8080"},
		},
		{
			netw: "tcp4",
			addr: []string{"127.0.0.1:8080", ":8080"},
		},
		{
			netw: "tcp6",
			addr: []string{"[::1]:8080", ":8080"},
		},
	}
	// START OMIT
	for _, tc := range tt {
		b.Run(tc.netw, func(b *testing.B) {
			for _, a := range tc.addr {
				b.Run(a, func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						ln, err := net.Listen(tc.netw, a)
						if err != nil {
							b.Fatalf("Listen error: %v", err)
						}
						ln.Close()
					}
				})
			}
		})
	}
	// END OMIT
}
