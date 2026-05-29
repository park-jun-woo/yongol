//ff:func feature=rule type=test control=sequence
//ff:what populateDDLIndexes — indexes 없을 때 빈 셋 key 는 여전히 존재

package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// TestPopulateDDLIndexes_NoIndexes — empty set key still exists.
func TestPopulateDDLIndexes_NoIndexes(t *testing.T) {
	g := newGround()
	populateDDLIndexes(g, ddl.Table{Name: "empty"})
	if _, ok := g.Lookup["DDL.index.empty"]; !ok {
		t.Errorf("DDL.index.empty key should exist (possibly empty)")
	}
}
