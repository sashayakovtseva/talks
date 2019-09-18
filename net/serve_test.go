package main

import (
	"context"
	"net"
	"net/http"
	"testing"
	"time"
)

func BenchmarkServe(b *testing.B) {
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
			addr: []string{"[::1]:8080", "[::]:8080"},
		},
	}
	for _, tc := range tt {
		b.Run(tc.netw, func(b *testing.B) {
			for _, a := range tc.addr {
				// START1 OMIT
				ln, err := net.Listen(tc.netw, a)
				if err != nil {
					b.Fatalf("Listen error: %v", err)
				}
				s := http.Server{
					Handler: http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
						w.Write([]byte("Hello!"))
					}),
				}
				// END1 OMIT

				go s.Serve(ln)
				time.Sleep(time.Second)

				// START2 OMIT
				b.Run(a, func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						resp, err := http.DefaultClient.Get("http://" + a)
						if err != nil {
							b.Fatal(err)
						}
						_ = resp.Body.Close()
					}
				})
				// END2 OMIT
				s.Shutdown(context.Background())
				ln.Close()
			}
		})
	}
}
