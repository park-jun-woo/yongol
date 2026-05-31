//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what xss60BuildTableMap/xss60ResolveVarModel/xss60FindMsgStruct 단위 검증
package ssac

import (
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

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
