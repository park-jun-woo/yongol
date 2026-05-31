//ff:func feature=validate type=test control=sequence topic=ddl-rego
//ff:what Run/helper test — XDP 규칙 일괄 실행 + buildDDLTableSet/ColumnIndex 검증
package ddl_rego

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
