//ff:func feature=rule type=test control=sequence
//ff:what populateOpenAPIParams — nil doc 허용 회귀

package ground

import (
	"testing"
)

// TestPopulateOpenAPIParams_NilDoc: nil doc must be tolerated.
func TestPopulateOpenAPIParams_NilDoc(t *testing.T) {
	g := newGround()
	populateOpenAPIParams(g, newMinimalFullstack())
	if len(g.Lookup) != 0 {
		t.Errorf("expected 0 keys, got %d", len(g.Lookup))
	}
}
