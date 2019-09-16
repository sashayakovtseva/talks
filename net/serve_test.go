package main

import (
	"context"
	"net"
	"net/http"
	"testing"
	"time"
)

func BenchmarkListenAndServeSingle(b *testing.B) {
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
		ln, err := net.Listen(tc.netw, tc.addr)
		if err != nil {
			b.Fatalf("Listen error: %v", err)
		}

		s := http.Server{
			Handler: http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.Write([]byte("Hello!"))
			}),
		}

		go s.Serve(ln)
		time.Sleep(time.Second)

		b.ResetTimer()
		b.Run(tc.netw, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				resp, err := http.DefaultClient.Get("http://" + tc.addr)
				if err != nil {
					b.Fatal(err)
				}
				_ = resp.Body.Close()
			}
		})
		b.StopTimer()
		s.Shutdown(context.Background())
		ln.Close()
	}
}

func BenchmarkListenAndServeAll(b *testing.B) {
	tt := []struct {
		netw string
		addr string
	}{
		{
			netw: "tcp",
			addr: "0.0.0.0:8080",
		},
		{
			netw: "tcp4",
			addr: "0.0.0.0:8080",
		},
		{
			netw: "tcp6",
			addr: "[::]:8080",
		},
	}
	for _, tc := range tt {
		ln, err := net.Listen(tc.netw, tc.addr)
		if err != nil {
			b.Fatalf("Listen error: %v", err)
		}

		s := http.Server{
			Handler: http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.Write([]byte("Hello!"))
			}),
		}

		go s.Serve(ln)
		time.Sleep(time.Second)

		b.ResetTimer()
		b.Run(tc.netw, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				resp, err := http.DefaultClient.Get("http://" + tc.addr)
				if err != nil {
					b.Fatal(err)
				}
				_ = resp.Body.Close()
			}
		})
		b.StopTimer()
		s.Shutdown(context.Background())
		ln.Close()
	}
}
