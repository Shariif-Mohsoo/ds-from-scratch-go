package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	opts := TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: NOPHandShakeFunc,
		Decoder:       DefaultDecoder{},
	}
	// The TCP address/port our test transport will listen on.
	// ":4000" means listen on port 4000 on all available network interfaces.
	// listenAddr := ":4000"

	// Create a new TCPTransport instance that will listen on the given address.
	tr := NewTCPTransport(opts)

	// Check that the transport's listenAddress matches the one we passed in.
	// If they don't match, the test fails.
	// assert.Equal(t, tr.l istenAddress, listenAddr)

	assert.Equal(t, tr.ListenAddr, ":3000")

	// Start the TCPTransport's listener.
	// assert.Nil checks that no error was returned — if there was an error,
	// the test will fail.
	// This makes sure that our transport can successfully start listening.
	assert.Nil(t, tr.ListenAndAccept())

	// Keep the test running indefinitely so the TCP listener stays alive.
	// In real tests, we'd probably connect to it from another transport to test actual communication.
	// select {}
}
