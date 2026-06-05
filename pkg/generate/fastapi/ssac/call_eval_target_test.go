//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestCallEvalTarget — callEvalTarget @call/@eval 대상 (package, function) 추출·빈값 분기 검증
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestCallEvalTarget(t *testing.T) {
	cases := []struct {
		name    string
		op      ir.Op
		wantPkg string
		wantFn  string
	}{
		{
			name:    "call op",
			op:      ir.Op{Kind: ir.OpCall, Call: &ir.CallOp{Package: "billing", Function: "HoldEscrow"}},
			wantPkg: "billing",
			wantFn:  "HoldEscrow",
		},
		{
			name:    "eval op",
			op:      ir.Op{Kind: ir.OpEval, Eval: &ir.EvalOp{Package: "auth", Function: "IsExpired"}},
			wantPkg: "auth",
			wantFn:  "IsExpired",
		},
		{
			name:    "call op nil payload",
			op:      ir.Op{Kind: ir.OpCall, Call: nil},
			wantPkg: "",
			wantFn:  "",
		},
		{
			name:    "eval op nil payload",
			op:      ir.Op{Kind: ir.OpEval, Eval: nil},
			wantPkg: "",
			wantFn:  "",
		},
		{
			name:    "other kind",
			op:      ir.Op{Kind: ir.OpAuth},
			wantPkg: "",
			wantFn:  "",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			pkg, fn := callEvalTarget(c.op)
			if pkg != c.wantPkg || fn != c.wantFn {
				t.Errorf("callEvalTarget() = (%q, %q), want (%q, %q)", pkg, fn, c.wantPkg, c.wantFn)
			}
		})
	}
}
