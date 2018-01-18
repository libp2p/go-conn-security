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

type InsecureConn struct {
	net.Conn
	local peer.ID
}

func (ic *InsecureConn) LocalPeer() peer.ID {
	return ic.local
}

func (ic *InsecureConn) RemotePeer() peer.ID {
	return ""
}

func (ic *InsecureConn) RemotePublicKey() ci.PubKey {
	return nil
}

func (ic *InsecureConn) LocalPrivateKey() ci.PrivKey {
	return nil
}

func (t InsecureTransport) Secure(ctx context.Context, insecure net.Conn, server bool) (Conn, error) {
	return &InsecureConn{
		Conn:  insecure,
		local: peer.ID(t),
	}, nil
}
