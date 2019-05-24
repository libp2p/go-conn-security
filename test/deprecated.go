// Deprecated: use github.com/libp2p/go-libp2p-testing/suites/sec instead.
package ss_test

import (
	"testing"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/sec"

	tsec "github.com/libp2p/go-libp2p-testing/suites/sec"
)

// Deprecated: use github.com/libp2p/go-libp2p-testing/suites/sec.Subtests instead.
var Subtests = tsec.Subtests

// Deprecated: use github.com/libp2p/go-libp2p-testing/suites/sec.TestMessage instead.
// Warning: it's impossible to alias a var in go. Writes to this var would have no longer
// have any effect, so it has been commented out to induce breakage for added safety.
// var TestMessage = tsec.TestMessage

// Deprecated: use github.com/libp2p/go-libp2p-testing/suites/sec.TestStreamLen instead.
// Warning: it's impossible to alias a var in go. Writes to this var would have no longer
// have any effect, so it has been commented out to induce breakage for added safety.
// var TestStreamLen = tsec.TestStreamLen

// Deprecated: use github.com/libp2p/go-libp2p-testing/suites/sec.TestSeed instead.
// Warning: it's impossible to alias a var in go. Writes to this var would have no longer
// have any effect, so it has been commented out to induce breakage for added safety.
// var TestSeed = tsec.TestSeed

// Deprecated: use github.com/libp2p/go-libp2p-testing/suites/sec.SubtestAll instead.
func SubtestAll(t *testing.T, at, bt sec.SecureTransport, ap, bp peer.ID) {
	tsec.SubtestAll(t, at, bt, ap, bp)
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/suites/sec.SubtestKeys instead.
func SubtestKeys(t *testing.T, at, bt sec.SecureTransport, ap, bp peer.ID) {
	tsec.SubtestKeys(t, at, bt, ap, bp)
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/suites/sec.SubtestWrongPeer instead.
func SubtestWrongPeer(t *testing.T, at, bt sec.SecureTransport, ap, bp peer.ID) {
	tsec.SubtestWrongPeer(t, at, bt, ap, bp)
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/suites/sec.SubtestCancelHandshakeOutbound instead.
func SubtestCancelHandshakeOutbound(t *testing.T, at, bt sec.SecureTransport, ap, bp peer.ID) {
	tsec.SubtestCancelHandshakeOutbound(t, at, bt, ap, bp)
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/suites/sec.SubtestCancelHandshakeInbound instead.
func SubtestCancelHandshakeInbound(t *testing.T, at, bt sec.SecureTransport, ap, bp peer.ID) {
	tsec.SubtestCancelHandshakeInbound(t, at, bt, ap, bp)
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/suites/sec.SubtestStream instead.
func SubtestStream(t *testing.T, at, bt sec.SecureTransport, ap, bp peer.ID) {
	tsec.SubtestStream(t, at, bt, ap, bp)
}
