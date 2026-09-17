package comms

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
)

type HTTPSChannel struct {
	address string
	client  *http.Client
	conn    net.Conn
}

func NewHTTPSChannel(addr string) *HTTPSChannel {
	return &HTTPSChannel{
		address: addr,
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ponytail: skip for self-signed
			},
		},
	}
}

func (c *HTTPSChannel) Connect() error {
	// We don't hold a persistent TCP conn in HTTPS but we verify the server is up.
	resp, err := c.client.Get("https://" + c.address)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}

func (c *HTTPSChannel) Send(data []byte) error {
	// ponytail: HTTPS a real C2 would use POST requests for data transfer.
	// For the sake of the interface, we simulate a POST.
	// In a real implementation, we'd implement a specialized HTTP transport.
	return fmt.Errorf("HTTPS Send requires specialized HTTP transport")
}

func (c *HTTPSChannel) Receive() ([]byte, error) {
	return nil, fmt.Errorf("HTTPS Receive requires specialized HTTP transport")
}

func (c *HTTPSChannel) Close() error {
	return nil
}
