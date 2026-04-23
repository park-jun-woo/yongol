//ff:func feature=rule type=test control=sequence
//ff:what populateDDLVarchar — VarcharLen 부재 시 Types 미기록

package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// TestPopulateDDLVarchar_Empty ensures no keys are written when VarcharLen is
// empty.
func TestPopulateDDLVarchar_Empty(t *testing.T) {
	g := newGround()
	populateDDLVarchar(g, ddl.Table{Name: "x"})
	if len(g.Types) != 0 {
		t.Errorf("expected no Types entries, got %d", len(g.Types))
	}
}
