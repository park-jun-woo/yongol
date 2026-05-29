//ff:func feature=rule type=test control=sequence
//ff:what TestVarDeclared_Declared_NoViolation — Ground.Vars 에 claim 이 있으면 위반 없음

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
