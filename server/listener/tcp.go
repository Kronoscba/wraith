package listener

import (
	"fmt"
	"net"
)

type TCPListener struct {
	address string
	ln      net.Listener
}

func NewTCPListener(addr string) *TCPListener {
	return &TCPListener{address: addr}
}

func (l *TCPListener) Start(handler func(net.Conn)) error {
	var err error
	l.ln, err = net.Listen("tcp", l.address)
	if err != nil {
		return fmt.Errorf("failed to start TCP listener: %w", err)
	}

	go func() {
		for {
			conn, err := l.ln.Accept()
			if err != nil {
				if l.ln == nil {
					return
				}
				continue
			}
			go handler(conn)
		}
	}()

	return nil
}

func (l *TCPListener) Stop() error {
	if l.ln != nil {
		return l.ln.Close()
	}
	return nil
}

func (l *TCPListener) GetAddress() net.Addr {
	if l.ln != nil {
		return l.ln.Addr()
	}
	return nil
}
