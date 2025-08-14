package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer represents a remote node (another computer or program)
// that we are connected to using a TCP connection.
type TCPPeer struct {
	// conn is the actual network connection to the peer (like a phone line between two people)
	conn net.Conn

	// outbound tells us if we started the connection (true) or if they connected to us (false)
	// outbound == true  → we called them (client)
	// outbound == false → they called us (server)
	outbound bool
}

// NewTCPPeer creates and returns a pointer to a new TCPPeer instance.
// Parameters:
// - conn: the TCP connection to this peer
// - outbound: whether this connection was started by us (true) or by them (false)
func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

// TCPTransport handles listening for new TCP connections
// and keeping track of connected peers.
type TCPTransport struct {
	// listenAddress is where this transport listens for incoming TCP connections
	// Example: ":8080" means listen on port 8080 on all interfaces
	listenAddress string

	// listener is the TCP listener that accepts new connections
	listener net.Listener

	// mu (mutex) is used to lock and protect 'peers' map from concurrent access
	mu sync.Mutex

	// peers stores all connected peers, using their network address as the key
	peers map[net.Addr]Peer
}

// NewTCPTransport creates and returns a pointer to a TCPTransport instance.
// Parameters:
// - listenAddr: the TCP address/port where we will listen for new connections
func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddr,
	}
}

// ListenAndAccept starts listening for incoming TCP connections
// and begins accepting them in a separate loop.
//
// Returns:
// - error: if something goes wrong while trying to start listening.
//
// How it works:
//  1. It calls net.Listen() to start a TCP listener on the given address.
//  2. If successful, it runs startAcceptLoop() in a goroutine
//     so it can handle new connections without blocking other code.
func (t *TCPTransport) ListenAndAccept() error {
	var err error

	// Start listening for TCP connections on t.listenAddress (e.g., ":8080")
	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		// If listening fails (e.g., port already in use), return the error
		return err
	}

	// Run the accept loop in a separate goroutine
	// so the program can keep doing other things while waiting for new peers
	go t.startAcceptLoop()

	// No errors — return nil
	return nil
}

// startAcceptLoop runs in a loop, constantly waiting for new TCP connections.
// When a new peer connects, it passes the connection to handleConn() in a new goroutine
// so that multiple peers can be handled at the same time without blocking each other.
func (t *TCPTransport) startAcceptLoop() {
	for {
		// Wait for a new incoming TCP connection.
		// This will "pause" here until a peer connects.
		conn, err := t.listener.Accept()
		if err != nil {
			// If something goes wrong (e.g., network error), log it and continue.
			fmt.Printf("TCP accept error: %s\n", err)
		}

		// Once we have a new connection, handle it in a separate goroutine.
		// This allows the loop to immediately go back to listening for more peers.
		go t.handleConn(conn)
	}
}

// handleConn is called whenever we get a new TCP connection.
// It wraps the raw connection into a TCPPeer object and logs the connection.
//
// Parameters:
// - conn: the raw TCP connection to the peer
func (t *TCPTransport) handleConn(conn net.Conn) {
	// Create a new TCPPeer object for this connection.
	// Here we are setting 'outbound' to true,
	// but in a real system we might check whether this connection
	// was initiated by us or by the remote peer.
	peer := NewTCPPeer(conn, true)

	// Print out the details of the new peer for debugging/logging purposes.
	// %+v shows field names and their values.
	fmt.Printf("new incoming connection %+v\n", peer)
}
