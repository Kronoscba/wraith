package listener

import (
	"fmt"
	"net"
	"net/http"
)

type HTTPListener struct {
	address string
	server  *http.Server
}

func NewHTTPListener(addr string) *HTTPListener {
	return &HTTPListener{address: addr}
}

func (l *HTTPListener) Start(handler func(net.Conn)) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Wraith-ID") != "" {
			hj, ok := w.(http.Hijacker)
			if !ok {
				return
			}
			conn, _, err := hj.Hijack()
			if err != nil {
				return
			}
			handler(conn)
			return
		}
		fmt.Fprintf(w, "Welcome to the public site")
	})

	l.server = &http.Server{
		Addr:    l.address,
		Handler: mux,
	}

	go func() {
		if err := l.server.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Printf("HTTP server error: %v\n", err)
		}
	}()

	return nil
}

func (l *HTTPListener) Stop() error {
	if l.server != nil {
		return l.server.Close()
	}
	return nil
}

func (l *HTTPListener) GetAddress() net.Addr {
	// This is a simplified version. 
	// In a real HTTP server, you'd return the listener's address.
	return nil 
}
