package insecure

import (
	"github.com/libp2p/go-libp2p-core/peer"
	moved "github.com/libp2p/go-libp2p-core/sec/insecure"
)

// Deprecated: Use github.com/libp2p/go-libp2p-core/sec/insecure.ID instead.
const ID = moved.ID

// Deprecated: Use github.com/libp2p/go-libp2p-core/sec/insecure.Transport instead.
type Transport = moved.Transport

// Deprecated: Use github.com/libp2p/go-libp2p-core/sec/insecure.New instead.
func New(id peer.ID) *moved.Transport {
	return moved.New(id)
}

// Deprecated: Use github.com/libp2p/go-libp2p-core/sec/insecure.Conn instead.
type Conn = moved.Conn
