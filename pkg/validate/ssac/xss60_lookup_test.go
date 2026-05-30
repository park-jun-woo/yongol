//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-structural
//ff:what xss60BuildTableMap/xss60ResolveVarModel/xss60FindMsgStruct 단위 검증

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
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

func TestXss60ResolveVarModel(t *testing.T) {
	fn := parsessac.ServiceFunc{Sequences: []parsessac.Sequence{
		{Type: "get", Result: &parsessac.Result{Var: "course", Type: "Course"}},
		{Type: "post", Result: &parsessac.Result{Var: "res", Type: "Reservation"}},
		{Type: "auth"}, // non-CRUD, ignored
	}}
	if got := xss60ResolveVarModel("course", fn); got != "Course" {
		t.Errorf("course -> %q, want Course", got)
	}
	if got := xss60ResolveVarModel("res", fn); got != "Reservation" {
		t.Errorf("res -> %q, want Reservation", got)
	}
	if got := xss60ResolveVarModel("unknown", fn); got != "" {
		t.Errorf("unknown -> %q, want empty", got)
	}
}

func TestXss60FindMsgStruct(t *testing.T) {
	fn := parsessac.ServiceFunc{
		Param: &parsessac.ParamInfo{TypeName: "OnOrderCompletedMessage"},
		Structs: []parsessac.StructInfo{
			{Name: "Other"},
			{Name: "OnOrderCompletedMessage"},
		},
	}
	got := xss60FindMsgStruct(fn)
	if got == nil || got.Name != "OnOrderCompletedMessage" {
		t.Errorf("findMsgStruct = %v", got)
	}

	// no matching struct -> nil
	fn2 := parsessac.ServiceFunc{
		Param:   &parsessac.ParamInfo{TypeName: "Missing"},
		Structs: []parsessac.StructInfo{{Name: "Other"}},
	}
	if got := xss60FindMsgStruct(fn2); got != nil {
		t.Errorf("expected nil, got %v", got)
	}
}
