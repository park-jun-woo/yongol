//ff:func feature=validate type=test control=sequence topic=ssac-ddl
//ff:what TestSSaCDDLHelpers — unit tests for the pure ssac_ddl helper functions
package ssac_ddl

import (
	"testing"

	parserddl "github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestDDLColumnSet(t *testing.T) {
	// Ground lookup path.
	fs := &yongol.Fullstack{}
	fs.SetGround(&rule.Ground{Lookup: map[string]rule.StringSet{
		"DDL.column.users": {"id": true, "email": true},
	}})
	set, ok := ddlColumnSet(fs, "users")
	if !ok || !set["email"] {
		t.Errorf("ground column set = %v ok=%v", set, ok)
	}

	// Fallback path.
	fs2 := &yongol.Fullstack{DDLTables: []parserddl.Table{
		{Name: "orders", Columns: map[string]parserddl.Column{"id": {Name: "id"}}},
	}}
	set2, ok := ddlColumnSet(fs2, "orders")
	if !ok || !set2["id"] {
		t.Errorf("fallback column set = %v ok=%v", set2, ok)
	}
	// Unknown table → (nil,false).
	if _, ok := ddlColumnSet(fs2, "missing"); ok {
		t.Error("missing table should return ok=false")
	}
}
