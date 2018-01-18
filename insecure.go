package ss

import (
	"context"
	"net"

	ci "github.com/libp2p/go-libp2p-crypto"
	peer "github.com/libp2p/go-libp2p-peer"
)

// InsecureTransport is a no-op stream security transport. It provides no
// security and simply wraps connections in blank
type InsecureTransport peer.ID

func (t *InsecureTransport) LocalPeer() peer.ID {
	return peer.ID(*t)
}

func (t *InsecureTransport) LocalPrivateKey() ci.PrivKey {
	return nil
}

func (t *InsecureTransport) SecureInbound(ctx context.Context, insecure net.Conn) (Conn, error) {
	return &InsecureConn{
		Conn:  insecure,
		local: peer.ID(*t),
	}, nil
}

func (t *InsecureTransport) SecureOutbound(ctx context.Context, insecure net.Conn, p peer.ID) (Conn, error) {
	return &InsecureConn{
		Conn:   insecure,
		local:  peer.ID(*t),
		remote: p,
	}, nil
}

// InsecureConn is the connection type returned by the insecure transport.
type InsecureConn struct {
	net.Conn
	local  peer.ID
	remote peer.ID
}

// LocalPeer returns the local peer ID.
func (ic *InsecureConn) LocalPeer() peer.ID {
	return ic.local
}

// RemotePeer returns the remote peer ID if we initiated the dial. Otherwise, it
// returns "" (because this connection isn't actually secure).
func (ic *InsecureConn) RemotePeer() peer.ID {
	return ic.remote
}

func (ic *InsecureConn) RemotePublicKey() ci.PubKey {
	return nil
}

func (ic *InsecureConn) LocalPrivateKey() ci.PrivKey {
	return nil
}

var _ Transport = (*InsecureTransport)(nil)
var _ Conn = (*InsecureConn)(nil)
