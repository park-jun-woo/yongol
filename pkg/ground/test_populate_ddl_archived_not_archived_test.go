//ff:func feature=rule type=test control=sequence
//ff:what populateDDLArchived — Archived=false 시 table-level flag 미설정

package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// TestPopulateDDLArchived_NotArchived ensures no table-level flag is set when
// Archived=false.
func TestPopulateDDLArchived_NotArchived(t *testing.T) {
	tab := ddl.Table{Name: "plain", Archived: false}
	fs := newMinimalFullstack(withDDLTables(tab))
	g := newGround()

	populateDDLArchived(g, fs)

	if g.Flags["archived.plain"] {
		t.Errorf("archived.plain should not be set when Archived=false")
	}
}
