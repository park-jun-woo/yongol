//ff:func feature=validate type=test control=sequence topic=ddl-rego
//ff:what Run/helper test — XDP 규칙 일괄 실행 + buildDDLTableSet/ColumnIndex 검증
package ddl_rego

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
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
