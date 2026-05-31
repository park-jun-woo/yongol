//ff:func feature=gen-fastapi type=util control=selection
//ff:what callEvalTarget — Op 이 @call/@eval 이면 (package, function) 반환, 아니면 빈 값

package ssac

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// callEvalTarget returns the package and function name targeted by an @call or
// @eval op. Returns empty strings for any other op kind or nil payload.
func callEvalTarget(op ir.Op) (pkg, fn string) {
	switch op.Kind {
	case ir.OpCall:
		if op.Call != nil {
			return op.Call.Package, op.Call.Function
		}
	case ir.OpEval:
		if op.Eval != nil {
			return op.Eval.Package, op.Eval.Function
		}
	}
	return "", ""
}
