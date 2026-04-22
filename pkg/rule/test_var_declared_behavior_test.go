//ff:func feature=rule type=test control=sequence
//ff:what VarDeclared 실제 동작 경로 — Ground.Vars 기준 (false=선언됨), (true=미선언)

package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/pkg/toulmin"
)

func TestVarDeclared_Declared_NoViolation(t *testing.T) {
	ground := &Ground{
		Vars: StringSet{"user_id": true},
	}
	ctx := toulmin.NewContext()
	ctx.Set("ground", ground)
	ctx.Set("claim", "user_id")

	spec := &VarDeclaredSpec{
		BaseSpec: BaseSpec{Rule: "VD-1", Level: "ERROR", Message: "not declared"},
	}
	ok, ev := VarDeclared(ctx, toulmin.Specs{spec})
	if ok || ev != nil {
		t.Fatalf("VarDeclared(declared) = (%v, %v); want (false, nil)", ok, ev)
	}
}

func TestVarDeclared_Undeclared_Violation(t *testing.T) {
	ground := &Ground{
		Vars: StringSet{"known": true},
	}
	ctx := toulmin.NewContext()
	ctx.Set("ground", ground)
	ctx.Set("claim", "unknown")

	spec := &VarDeclaredSpec{
		BaseSpec: BaseSpec{Rule: "VD-2", Level: "ERROR", Message: "not declared"},
	}
	ok, ev := VarDeclared(ctx, toulmin.Specs{spec})
	if !ok {
		t.Fatalf("VarDeclared(undeclared) ok = false; want true")
	}
	e, okType := ev.(*Evidence)
	if !okType {
		t.Fatalf("VarDeclared(undeclared) evidence type = %T; want *Evidence", ev)
	}
	if e.Rule != "VD-2" || e.Ref != "unknown" || e.Level != "ERROR" {
		t.Fatalf("VarDeclared evidence = %+v; want Rule=VD-2 Ref=unknown Level=ERROR", e)
	}
}

func TestVarDeclared_NilSpecEntry_ReturnsFalse(t *testing.T) {
	ctx := toulmin.NewContext()
	var nilSpec *VarDeclaredSpec
	ok, ev := VarDeclared(ctx, toulmin.Specs{nilSpec})
	if ok || ev != nil {
		t.Fatalf("VarDeclared(nil spec) = (%v, %v); want (false, nil)", ok, ev)
	}
}
