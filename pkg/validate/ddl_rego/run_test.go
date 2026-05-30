//ff:func feature=validate type=test control=sequence topic=ddl-rego
//ff:what Run/helper test — XDP 규칙 일괄 실행 + buildDDLTableSet/ColumnIndex 검증

package ddl_rego

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildDDLTableSet(t *testing.T) {
	fs := &yongol.Fullstack{DDLTables: []ddl.Table{ddlTable("a"), ddlTable("b")}}

	// nil Ground -> fallback to fs.DDLTables.
	set := buildDDLTableSet(fs, nil)
	if !set["a"] || !set["b"] || set["c"] {
		t.Errorf("fallback set wrong: %v", set)
	}

	// Ground present -> prefer Ground.Lookup.
	fg := fsWithGroundTables([]string{"x"}, nil)
	fg.DDLTables = []ddl.Table{ddlTable("a")}
	set2 := buildDDLTableSet(fg, fg.Ground())
	if !set2["x"] || set2["a"] {
		t.Errorf("ground-preferred set wrong: %v", set2)
	}
}

func TestBuildDDLColumnIndex(t *testing.T) {
	fs := &yongol.Fullstack{DDLTables: []ddl.Table{ddlTable("users", "id", "email")}}
	idx := buildDDLColumnIndex(fs)
	if !idx["users"]["id"] || !idx["users"]["email"] {
		t.Errorf("column index missing cols: %v", idx)
	}
	if idx["users"]["ghost"] {
		t.Errorf("unexpected column present")
	}
}

func TestRun(t *testing.T) {
	// Empty Fullstack -> no diagnostics (each rule short-circuits).
	if d := Run(&yongol.Fullstack{}); len(d) != 0 {
		t.Errorf("empty fs Run should yield no diags, got %v", d)
	}

	// One XDP-31 violation aggregated through Run.
	fs := fsWithGroundTables([]string{"other"}, []rego.Policy{
		{File: "p.rego", Ownerships: []rego.OwnershipMapping{
			{Resource: "gig", Table: "gigs", SourceLine: 1},
		}},
	})
	d := Run(fs)
	if len(d) != 1 {
		t.Fatalf("expected 1 aggregated diag, got %d: %v", len(d), d)
	}
}
