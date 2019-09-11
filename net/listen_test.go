package main

import (
	"net"
	"testing"
)

func BenchmarkListenSingle(b *testing.B) {
	addr := "127.0.0.1:8080"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		ln, err := net.Listen("tcp", addr)
		b.StopTimer()
		if err != nil {
			b.Fatalf("Listen error: %v", err)
		}
		ln.Close()
	}
}

func BenchmarkListenAll(b *testing.B) {
	addr := ":8080"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		ln, err := net.Listen("tcp", addr)
		b.StopTimer()
		if err != nil {
			b.Fatalf("Listen error: %v", err)
		}
		ln.Close()
	}
}
