package listener

import "net"

// Listener defines the interface for all C2 communication listeners.
type Listener interface {
	// Start begins listening for agent connections.
	Start(handler func(net.Conn)) error
	// Stop ceases listening and closes all active connections.
	Stop() error
	// GetAddress returns the local address the listener is bound to.
	GetAddress() net.Addr
}
