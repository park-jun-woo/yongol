//ff:func feature=rule type=test control=sequence
//ff:what populateMiddleware — nil manifest 단락 회귀

package ground

import (
	"testing"
)

// TestPopulateMiddleware_NilManifest short-circuits safely.
func TestPopulateMiddleware_NilManifest(t *testing.T) {
	g := newGround()
	populateMiddleware(g, newMinimalFullstack())
	if len(g.Lookup) != 0 {
		t.Errorf("expected no keys, got %d", len(g.Lookup))
	}
}
