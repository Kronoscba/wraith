package comms

import (
	"fmt"
	"net"
)

type TCPChannel struct {
	address string
	conn    net.Conn
}

func NewTCPChannel(addr string) *TCPChannel {
	return &TCPChannel{address: addr}
}

func (c *TCPChannel) Connect() error {
	var err error
	c.conn, err = net.Dial("tcp", c.address)
	if err != nil {
		return fmt.Errorf("failed to connect to C2: %w", err)
	}
	return nil
}

func (c *TCPChannel) Send(data []byte) error {
	if c.conn == nil {
		return fmt.Errorf("not connected")
	}
	_, err := c.conn.Write(data)
	return err
}

func (c *TCPChannel) Receive() ([]byte, error) {
	if c.conn == nil {
		return nil, fmt.Errorf("not connected")
	}
	buf := make([]byte, 4096)
	n, err := c.conn.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf[:n], nil
}

func (c *TCPChannel) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}
