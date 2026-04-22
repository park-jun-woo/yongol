//ff:func feature=validate-contract type=util control=sequence topic=preserve-safety
//ff:what isCurrentUserAssertion — expr 가 `ctx.Value("currentUser").(T)` 타입 단언인지

package contract

import (
	"go/ast"
	"strconv"
)

// isCurrentUserAssertion reports whether expr is a TypeAssertExpr whose
// operand is `ctx.Value("currentUser")`. Only the AST shape is checked
// — the caller (PRV-11) is responsible for validating that the
// containing AssignStmt uses the comma-ok 2-tuple form.
//
// The ctx identifier name is intentionally pinned to "ctx" because
// every yongol-generated handler follows the
// `func ... (ctx context.Context, ...)` convention; other receivers
// should not appear in preserved handler bodies.
func isCurrentUserAssertion(expr ast.Expr) bool {
	ta, ok := expr.(*ast.TypeAssertExpr)
	if !ok || ta.Type == nil {
		return false
	}
	call, ok := ta.X.(*ast.CallExpr)
	if !ok || len(call.Args) != 1 {
		return false
	}
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok || sel.Sel == nil || sel.Sel.Name != "Value" {
		return false
	}
	recv, ok := sel.X.(*ast.Ident)
	if !ok || recv.Name != "ctx" {
		return false
	}
	lit, ok := call.Args[0].(*ast.BasicLit)
	if !ok {
		return false
	}
	val, err := strconv.Unquote(lit.Value)
	if err != nil {
		return false
	}
	return val == "currentUser"
}
