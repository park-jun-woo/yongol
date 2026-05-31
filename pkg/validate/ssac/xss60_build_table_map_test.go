//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what xss60BuildTableMap/xss60ResolveVarModel/xss60FindMsgStruct 단위 검증
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXss60BuildTableMap(t *testing.T) {
	fs := &yongol.Fullstack{DDLTables: []ddl.Table{{Name: "users"}, {Name: "courses"}}}
	m := xss60BuildTableMap(fs)
	if len(m) != 2 {
		t.Fatalf("len = %d, want 2", len(m))
	}
	if m["users"] == nil || m["users"].Name != "users" {
		t.Errorf("users entry = %v", m["users"])
	}
	if m["courses"] == nil || m["courses"].Name != "courses" {
		t.Errorf("courses entry = %v", m["courses"])
	}
	if _, ok := m["missing"]; ok {
		t.Errorf("missing should not be present")
	}
}
