//ff:func feature=rule type=test control=sequence
//ff:what populateOpenAPIResponseTypes — doc 부재 시 Types 미기록

package ground

import (
	"testing"
)

// TestPopulateOpenAPIResponseTypes_NilDoc — no panic when doc absent.
func TestPopulateOpenAPIResponseTypes_NilDoc(t *testing.T) {
	g := newGround()
	populateOpenAPIResponseTypes(g, newMinimalFullstack())
	if len(g.Types) != 0 {
		t.Errorf("expected no Types entries, got %d", len(g.Types))
	}
}
