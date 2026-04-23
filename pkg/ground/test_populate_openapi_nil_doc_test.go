//ff:func feature=rule type=test control=sequence
//ff:what populateOpenAPI — nil doc 단락 회귀

package ground

import (
	"testing"
)

// TestPopulateOpenAPI_NilDoc ensures the function short-circuits safely.
func TestPopulateOpenAPI_NilDoc(t *testing.T) {
	g := newGround()
	populateOpenAPI(g, newMinimalFullstack())

	if len(g.Lookup) != 0 {
		t.Errorf("expected empty Lookup when doc is nil, got %d keys", len(g.Lookup))
	}
}
