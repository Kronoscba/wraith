package listener

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
)

type HTTPSListener struct {
	address  string
	certFile string
	keyFile  string
	server   *http.Server
}

func NewHTTPSListener(addr, cert, key string) *HTTPSListener {
	return &HTTPSListener{
		address:  addr,
		certFile: cert,
		keyFile:  key,
	}
}

func (l *HTTPSListener) Start(handler func(net.Conn)) error {
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
		fmt.Fprintf(w, "Secure Gateway")
	})

	l.server = &http.Server{
		Addr:    l.address,
		Handler: mux,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	}

	go func() {
		if err := l.server.ListenAndServeTLS(l.certFile, l.keyFile); err != http.ErrServerClosed {
			fmt.Printf("HTTPS server error: %v\n", err)
		}
	}()

	return nil
}

func (l *HTTPSListener) Stop() error {
	if l.server != nil {
		return l.server.Close()
	}
	return nil
}

func (l *HTTPSListener) GetAddress() net.Addr {
	return nil
}
