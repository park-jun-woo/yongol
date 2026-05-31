//ff:func feature=validate type=test control=sequence topic=ssac-ddl
//ff:what TestSSaCDDLHelpers — unit tests for the pure ssac_ddl helper functions
package ssac_ddl

import (
	"testing"

	parserddl "github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestDDLTableSet(t *testing.T) {
	// Ground lookup path.
	fs := &yongol.Fullstack{}
	fs.SetGround(&rule.Ground{Lookup: map[string]rule.StringSet{
		"DDL.table": {"users": true, "courses": true},
	}})
	set := ddlTableSet(fs)
	if !set["users"] || !set["courses"] {
		t.Errorf("ground lookup set = %v", set)
	}

	// Fallback path (no Ground): build from DDLTables.
	fs2 := &yongol.Fullstack{DDLTables: []parserddl.Table{{Name: "orders"}}}
	set2 := ddlTableSet(fs2)
	if !set2["orders"] {
		t.Errorf("fallback set = %v", set2)
	}
}
