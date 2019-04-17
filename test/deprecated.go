package ss_test

import (
	"testing"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/sec"

	moved "github.com/libp2p/go-libp2p-core/sec/test"
)

// Deprecated: Use github.com/libp2p/go-libp2p-core/sec/test.Subtests instead.
var Subtests = moved.Subtests

// Deprecated: Use github.com/libp2p/go-libp2p-core/sec/test.TestMessage instead.
var TestMessage = moved.TestMessage

// Deprecated: Use github.com/libp2p/go-libp2p-core/sec/test.TestStreamLen instead.
// Warning: it's impossible to alias a var in go, so reads and writes to this variable may be inaccurate
// or not have the intended effect.
var TestStreamLen = moved.TestStreamLen

// Deprecated: Use github.com/libp2p/go-libp2p-core/sec/test.TestSeed instead.
// Warning: it's impossible to alias a var in go, so reads and writes to this variable may be inaccurate
// or not have the intended effect.
var TestSeed = moved.TestSeed

// Deprecated: Use github.com/libp2p/go-libp2p-core/sec/test.SubtestAll instead.
func SubtestAll(t *testing.T, at, bt sec.SecureTransport, ap, bp peer.ID) {
	moved.SubtestAll(t, at, bt, ap, bp)
}

// Deprecated: Use github.com/libp2p/go-libp2p-core/sec/test.SubtestKeys instead.
func SubtestKeys(t *testing.T, at, bt sec.SecureTransport, ap, bp peer.ID) {
	moved.SubtestKeys(t, at, bt, ap, bp)
}

// Deprecated: Use github.com/libp2p/go-libp2p-core/sec/test.SubtestWrongPeer instead.
func SubtestWrongPeer(t *testing.T, at, bt sec.SecureTransport, ap, bp peer.ID) {
	moved.SubtestWrongPeer(t, at, bt, ap, bp)
}

// Deprecated: Use github.com/libp2p/go-libp2p-core/sec/test.SubtestCancelHandshakeOutbound instead.
func SubtestCancelHandshakeOutbound(t *testing.T, at, bt sec.SecureTransport, ap, bp peer.ID) {
	moved.SubtestCancelHandshakeOutbound(t, at, bt, ap, bp)
}

// Deprecated: Use github.com/libp2p/go-libp2p-core/sec/test.SubtestCancelHandshakeInbound instead.
func SubtestCancelHandshakeInbound(t *testing.T, at, bt sec.SecureTransport, ap, bp peer.ID) {
	moved.SubtestCancelHandshakeInbound(t, at, bt, ap, bp)
}

// Deprecated: Use github.com/libp2p/go-libp2p-core/sec/test.SubtestStream instead.
func SubtestStream(t *testing.T, at, bt sec.SecureTransport, ap, bp peer.ID) {
	moved.SubtestStream(t, at, bt, ap, bp)
}
