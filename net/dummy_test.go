package main

import (
	"net"
	"testing"
)

func BenchmarkDummyListen(b *testing.B) {
	tt := []string{
		"127.0.0.1:8080",
		":8080",
	}

	for _, tc := range tt {
		b.Run(tc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ln, err := net.Listen("tcp", tc)
				if err != nil {
					b.Fatalf("Listen error: %v", err)
				}
				ln.Close()
			}
		})
	}
}
