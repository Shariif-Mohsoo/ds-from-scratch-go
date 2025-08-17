package p2p

// ErrInvalidHandshake is returned if the handshake between
// the local and remote node could not be established.
// var ErrInvalidHandshake = errors.New("Invalid Handshake")

type HandshakeFunc func(Peer) error

func NOPHandShakeFunc(Peer) error { return nil }
