package comms

// Channel defines the interface for all agent communication channels.
type Channel interface {
	// Connect establishes a connection to the C2 server.
	Connect() error
	// Send transmits a payload to the C2 server.
	Send(data []byte) error
	// Receive waits for and retrieves a payload from the C2 server.
	Receive() ([]byte, error)
	// Close terminates the connection to the C2 server.
	Close() error
}
