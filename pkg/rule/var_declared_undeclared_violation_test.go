//ff:func feature=rule type=test control=sequence
//ff:what TestVarDeclared_Undeclared_Violation — Ground.Vars 에 claim 이 없으면 Evidence 방출

package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/pkg/toulmin"
)

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
