package listener

import (
	"fmt"
	"net"
	"time"
)

type DNSListener struct {
	address string
	conn    *net.UDPConn
}

func NewDNSListener(addr string) *DNSListener {
	return &DNSListener{address: addr}
}

func (l *DNSListener) Start(handler func(net.Conn)) error {
	addr, err := net.ResolveUDPAddr("udp", l.address)
	if err != nil {
		return err
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return err
	}
	l.conn = conn

	go func() {
		buf := make([]byte, 512)
		for {
			n, remoteAddr, err := l.conn.ReadFromUDP(buf)
			if err != nil {
				continue
			}
			
			// ponytail: DNS Tunneling. 
			// We extract the query and wrap it in a dummy net.Conn 
			// so the handler doesn't know it's UDP/DNS.
			go func(data []byte, addr *net.UDPAddr) {
				// In a real DNS C2, we'd parse the DNS packet here.
				// We use a custom wrapper that implements net.Conn over UDP.
				// For now, we simulate a connection.
				handler(newUDPConnWrapper(l.conn, addr, data))
			}(append([]byte(nil), buf[:n]...), remoteAddr)
		}
	}()

	return nil
}

func (l *DNSListener) Stop() error {
	if l.conn != nil {
		return l.conn.Close()
	}
	return nil
}

func (l *DNSListener) GetAddress() net.Addr {
	if l.conn != nil {
		return l.conn.LocalAddr()
	}
	return nil
}

// UDPConnWrapper mimics net.Conn over a UDP socket.
type UDPConnWrapper struct {
	conn    *net.UDPConn
	remote  *net.UDPAddr
	buffer  []byte
	pos     int
}

func newUDPConnWrapper(c *net.UDPConn, r *net.UDPAddr, initial []byte) net.Conn {
	return &UDPConnWrapper{conn: c, remote: r, buffer: initial}
}

func (u *UDPConnWrapper) Read(b []byte) (n int, err error) {
	if u.pos >= len(u.buffer) {
		// In a real DNS C2, this would wait for the next DNS query.
		return 0, fmt.Errorf("no more data in DNS packet")
	}
	n = copy(b, u.buffer[u.pos:])
	u.pos += n
	return n, nil
}

func (u *UDPConnWrapper) Write(b []byte) (n int, err error) {
	// Respond with a DNS TXT record or similar.
	return u.conn.WriteToUDP(b, u.remote)
}

func (u *UDPConnWrapper) Close() error { return nil }
func (u *UDPConnWrapper) LocalAddr() net.Addr { return u.conn.LocalAddr() }
func (u *UDPConnWrapper) RemoteAddr() net.Addr { return u.remote }
func (u *UDPConnWrapper) SetDeadline(t time.Time) error { return nil }
func (u *UDPConnWrapper) SetReadDeadline(t time.Time) error { return nil }
func (u *UDPConnWrapper) SetWriteDeadline(t time.Time) error { return nil }
