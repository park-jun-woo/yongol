//ff:func feature=rule type=test control=sequence dimension=1
//ff:what populateDDL — 테이블, 컬럼, default 값이 Lookup/Flags/Types에 등록

package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// TestPopulateDDL_BasicColumnsAndDefaults checks that populateDDL registers
// table names, per-table column sets, and default annotations.
func TestPopulateDDL_BasicColumnsAndDefaults(t *testing.T) {
	tab := ddl.Table{
		Name: "users",
		Columns: map[string]ddl.Column{
			"id":    {Name: "id", RawType: "BIGINT"},
			"email": {Name: "email", RawType: "VARCHAR(255)", HasDefault: true, DefaultLiteral: ""},
		},
	}
	fs := newMinimalFullstack(withDDLTables(tab))
	g := newGround()

	populateDDL(g, fs)

	tables := g.Lookup["DDL.table"]
	if !tables["users"] {
		t.Fatalf("DDL.table missing users: %v", tables)
	}

	cols := g.Lookup["DDL.column.users"]
	if !cols["id"] || !cols["email"] {
		t.Fatalf("DDL.column.users = %v, want id,email", cols)
	}

	if !g.Flags["DDL.default.users.email"] {
		t.Errorf("DDL.default.users.email flag missing")
	}
	if _, ok := g.Types["DDL.default.value.users.email"]; !ok {
		t.Errorf("DDL.default.value.users.email missing")
	}
}
