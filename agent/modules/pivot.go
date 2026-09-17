package modules

import (
	"fmt"
	"net"
)

type Peer struct {
	ID   string
	Addr string
}

// Pivot establishes a connection to another agent.
func Pivot(peerAddr string) (net.Conn, error) {
	// ponytail: basic TCP proxying.
	// In a real C2, this would be a multi-hop encrypted tunnel.
	conn, err := net.Dial("tcp", peerAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to pivot to %s: %w", peerAddr, err)
	}
	return conn, nil
}
