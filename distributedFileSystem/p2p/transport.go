package p2p

// Peer represents a remote node in the network.
// In Go, an interface is just a "contract" — it defines behavior
// that other types must implement to be considered a Peer.
//
// Right now, this interface is empty, which means ANY type
// automatically satisfies it. (Later, we can add methods like `Close()` or `Send()`.)
type Peer interface {
}

// Transport is an interface for anything that handles communication
// between nodes in the network.
//
// Examples of transports could be:
// - TCPTransport (over regular TCP sockets)
// - UDPTransport (over UDP)
// - WebSocketTransport (over WebSockets)
//
// By using an interface, we make the code flexible so we can swap out
// the communication method without changing the rest of the system.
type Transport interface {
	// ListenAndAccept should start listening for new peer connections
	// and accept them as they arrive.
	ListenAndAccept() error
}
