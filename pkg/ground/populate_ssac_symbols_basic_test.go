//ff:func feature=rule type=test control=sequence
//ff:what populateSSaCSymbols — DDL row → Struct.<Model>.<PascalField> 타입 등록

package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// TestPopulateSSaCSymbols_StructTypes verifies DDL table rows are registered
// as Struct.<Model>.<PascalField> = <GoType>. Model is singular PascalCase of
// the table name; field is PascalCase of the column.
func TestPopulateSSaCSymbols_StructTypes(t *testing.T) {
	tab := ddl.Table{
		Name: "users",
		Columns: map[string]ddl.Column{
			"id":         {Name: "id", RawType: "BIGINT"},
			"created_at": {Name: "created_at", RawType: "TIMESTAMPTZ"},
		},
	}
	fs := newMinimalFullstack(withDDLTables(tab))
	g := newGround()

	populateSSaCSymbols(g, fs)

	// "users" (plural) → singular "User" PascalCase
	if got := g.Types["Struct.User.ID"]; got != "int64" {
		t.Errorf("Struct.User.ID = %q, want int64", got)
	}
	if got := g.Types["Struct.User.CreatedAt"]; got != "time.Time" {
		t.Errorf("Struct.User.CreatedAt = %q, want time.Time", got)
	}
}
