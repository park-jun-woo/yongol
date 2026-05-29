//ff:func feature=rule type=test control=sequence
//ff:what populateDDLVarchar — VARCHAR(N) 길이가 Types["DDL.varchar..."]로 등록

package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// TestPopulateDDLVarchar_MaxLen verifies each column's varchar length is
// registered under DDL.varchar.<table>.<col> as a decimal string.
func TestPopulateDDLVarchar_MaxLen(t *testing.T) {
	tab := ddl.Table{
		Name: "users",
		Columns: map[string]ddl.Column{
			"email": {Name: "email", RawType: "VARCHAR(255)", VarcharLen: 255},
			"name":  {Name: "name", RawType: "VARCHAR(50)", VarcharLen: 50},
		},
	}
	g := newGround()
	populateDDLVarchar(g, tab)

	if got := g.Types["DDL.varchar.users.email"]; got != "255" {
		t.Errorf("DDL.varchar.users.email = %q, want 255", got)
	}
	if got := g.Types["DDL.varchar.users.name"]; got != "50" {
		t.Errorf("DDL.varchar.users.name = %q, want 50", got)
	}
}
