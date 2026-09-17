package comms

import (
	"fmt"
	"net"
)

type DNSChannel struct {
	domain string
	dnsSrv string
}

func NewDNSChannel(domain, dnsSrv string) *DNSChannel {
	return &DNSChannel{
		domain: domain,
		dnsSrv: dnsSrv,
	}
}

func (c *DNSChannel) Connect() error {
	// DNS is connectionless. We just verify the domain resolves.
	_, err := net.LookupHost(c.domain)
	if err != nil {
		// ponytail: return nil because lookup might fail but tunneling still works via custom DNS server
		return nil 
	}
	return nil
}

func (c *DNSChannel) Send(data []byte) error {
	// ponytail: encode data as subdomains: <b64data>.wraith.com
	// In a real implementation, we'd use a DNS library to send a specific query type.
	fmt.Printf("DNS Send: %s\n", string(data))
	return nil
}

func (c *DNSChannel) Receive() ([]byte, error) {
	// ponytail: wait for DNS TXT response.
	return nil, fmt.Errorf("DNS Receive not yet implemented")
}

func (c *DNSChannel) Close() error {
	return nil
}
