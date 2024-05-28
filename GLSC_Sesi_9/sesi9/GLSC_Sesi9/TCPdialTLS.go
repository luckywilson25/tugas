package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"strings"
	"time"
)

func main() {
	var address string

	
	fmt.Println("Input Web Address (Example: Google.com)")
	fmt.Print("> ")
	fmt.Scanln(&address)

	address = strings.TrimPrefix(address, "https://")
	address = strings.TrimPrefix(address, "http://")

	if !strings.Contains(address, ":") {
		address = address + ":443"
	}

	// Define the dialer with a timeout
	dialer := &net.Dialer{
		Timeout: 5 * time.Second,
	}

	// Connect to the HTTPS website
	conn, err := tls.DialWithDialer(dialer, "tcp", address, &tls.Config{
		InsecureSkipVerify: true, // Skip verification for demonstration purposes
	})
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	// Get the connection state
	state := conn.ConnectionState()

	// Print the TLS version
	fmt.Println("TLS Version:", tlsVersion(state.Version))

	// Print the CipherSuite name
	fmt.Println("Cipher Suite:", tls.CipherSuiteName(state.CipherSuite))

	// Print the Issuer Organization
	for _, cert := range state.PeerCertificates {
		issuer := cert.Issuer
		fmt.Println("Issuer Organization:", issuer.Organization)
		break // Only print the first certificate's issuer
	}
}

// Function to convert TLS version to human-readable format
func tlsVersion(version uint16) string {
	switch version {
	case tls.VersionTLS13:
		return "TLS 1.3"
	case tls.VersionTLS12:
		return "TLS 1.2"
	case tls.VersionTLS11:
		return "TLS 1.1"
	case tls.VersionTLS10:
		return "TLS 1.0"
	default:
		return "Unknown"
	}
}