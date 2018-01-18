package ss

import (
	"context"
	"net"

	ic "github.com/libp2p/go-libp2p-crypto"
	peer "github.com/libp2p/go-libp2p-peer"
	transport "github.com/libp2p/go-libp2p-transport"
)

// Transport is an interface that turns an unauthenticated, plain-text
// stream into an authenticated, encrypted stream.
type Transport interface {
	// Secure secures the passed insecure connection.
	Secure(ctx context.Context, insecure net.Conn, server bool) (Conn, error)

	// LocalPeer returns the local peer ID used by this security transport.
	LocalPeer() peer.ID

	// LocalPrivateKey returns the private key used by this security
	// transport.
	LocalPrivateKey() ic.PrivKey
}

// Conn provides the necessary functionality to wrap a connection
// and tunnel it through a secure channel via the provided ReadWriter.
type Conn interface {
	net.Conn
	transport.ConnSecurity
}
