//ff:func feature=rule type=test control=sequence
//ff:what populateManifest — nil manifest 단락 회귀 테스트

package ground

import (
	"testing"
)

// TestPopulateManifest_NilManifest: nil manifest must short-circuit.
func TestPopulateManifest_NilManifest(t *testing.T) {
	g := newGround()
	populateManifest(g, newMinimalFullstack())
	if len(g.Lookup) != 0 || len(g.Config) != 0 {
		t.Errorf("expected no keys, got Lookup=%d Config=%d", len(g.Lookup), len(g.Config))
	}
}
