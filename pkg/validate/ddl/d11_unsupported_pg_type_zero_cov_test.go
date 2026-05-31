//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what validate/ddl 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package ddl

import (
	"testing"

	pddl "github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestD11UnsupportedPgType_ZeroCov(t *testing.T) {
	cols := map[string]pddl.Column{
		"c": {Name: "c", RawType: "TIME WITH TIME ZONE"},
	}
	fs := &yongol.Fullstack{
		DDLTables: []pddl.Table{{Name: "t", File: "t.sql", Line: 1, Columns: cols}},
	}
	diags := d11UnsupportedPgType(fs)
	if len(diags) == 0 {
		t.Errorf("expected D-11 diag for unsupported type")
	}
	if got := d11UnsupportedPgType(nil); got != nil {
		t.Errorf("nil fs should give nil")
	}
}
