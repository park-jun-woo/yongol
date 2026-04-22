//ff:func feature=validate-contract type=util control=sequence topic=preserve-safety
//ff:what scanCallFromAssign — AssignStmt RHS 가 Scan 호출이면 반환

package contract

import "go/ast"

// scanCallFromAssign returns the underlying CallExpr when the
// assignment's RHS is a single `.Scan(...)` call on a row-like
// receiver. Otherwise returns nil so the caller can skip.
func scanCallFromAssign(as *ast.AssignStmt) *ast.CallExpr {
	if as == nil || len(as.Rhs) != 1 {
		return nil
	}
	call, ok := as.Rhs[0].(*ast.CallExpr)
	if !ok {
		return nil
	}
	if !isScanCall(call) {
		return nil
	}
	return call
}
