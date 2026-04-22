//ff:func feature=validate-contract type=util control=sequence topic=preserve-safety
//ff:what resourceAcquireFromAssign — AssignStmt 가 close 필요한 리소스 호출이면 var 이름 + CallExpr 반환

package contract

import "go/ast"

// resourceAcquireFromAssign inspects the RHS of as for a call that
// returns a closable resource (os.Open, http.Get, db.Query, ...). On
// match it returns the LHS variable name bound to the resource. The
// second return value is the CallExpr itself so the caller can report
// line numbers.
//
// LHS shapes accepted:
//   - single-value: `v := os.Open(...)` → returns "v".
//   - two-value:    `v, err := db.Query(...)` → returns "v".
//
// The blank identifier `_` as the resource target is skipped — the
// caller is explicitly throwing the handle away, so close-tracking
// is moot.
func resourceAcquireFromAssign(as *ast.AssignStmt) (string, *ast.CallExpr) {
	if as == nil || len(as.Rhs) != 1 {
		return "", nil
	}
	call, ok := as.Rhs[0].(*ast.CallExpr)
	if !ok || !isResourceAcquireCall(call) {
		return "", nil
	}
	if len(as.Lhs) == 0 {
		return "", nil
	}
	ident, ok := as.Lhs[0].(*ast.Ident)
	if !ok || ident.Name == "_" {
		return "", nil
	}
	return ident.Name, call
}
