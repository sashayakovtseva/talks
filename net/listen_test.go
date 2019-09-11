package main

import (
	"net"
	"testing"
)

func BenchmarkListenSingle(b *testing.B) {
	tt := []struct {
		netw string
		addr string
	}{
		{
			netw: "tcp",
			addr: "127.0.0.1:8080",
		},
		{
			netw: "tcp4",
			addr: "127.0.0.1:8080",
		},
		{
			netw: "tcp6",
			addr: "[::1]:8080",
		},
	}
	for _, tc := range tt {
		b.Run(tc.netw, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ln, err := net.Listen(tc.netw, tc.addr)
				if err != nil {
					b.Fatalf("Listen error: %v", err)
				}
				ln.Close()
			}
		})
	}
}

func BenchmarkListenAll(b *testing.B) {
	tt := []string{
		"tcp",
		"tcp4",
		"tcp6",
	}
	for _, tc := range tt {
		b.Run(tc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ln, err := net.Listen(tc, ":8080")
				if err != nil {
					b.Fatalf("Listen error: %v", err)
				}
				ln.Close()
			}
		})
	}
}
