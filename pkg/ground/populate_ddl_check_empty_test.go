//ff:func feature=rule type=test control=sequence
//ff:what populateDDLCheck — CheckEnums 부재 시 panic/extra key 없음

package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// TestPopulateDDLCheck_Empty verifies no panic/extra keys when no CHECK enum.
func TestPopulateDDLCheck_Empty(t *testing.T) {
	tab := ddl.Table{Name: "empty"}
	g := newGround()
	populateDDLCheck(g, tab)

	if _, ok := g.Lookup["DDL.check.empty.status"]; ok {
		t.Errorf("unexpected DDL.check key for empty CheckEnums")
	}
}
