package connsec

import (
	"context"
	"net"

	peer "github.com/libp2p/go-libp2p-peer"
	transport "github.com/libp2p/go-libp2p-transport"
)

// A Transport turns inbound and outbound unauthenticated,
// plain-text connections into authenticated, encrypted connections.
type Transport interface {
	// SecureInbound secures an inbound connection.
	SecureInbound(ctx context.Context, insecure net.Conn) (Conn, error)

	// SecureOutbound secures an outbound connection.
	SecureOutbound(ctx context.Context, insecure net.Conn, p peer.ID) (Conn, error)
}

// Conn is an authenticated, encrypted connection.
type Conn interface {
	net.Conn
	transport.ConnSecurity
}
