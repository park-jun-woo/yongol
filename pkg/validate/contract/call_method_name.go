//ff:func feature=validate-contract type=util control=sequence topic=preserve-safety
//ff:what callMethodName — CallExpr 가 SelectorExpr 라면 method 이름 반환 (receiver 호출만)

package contract

import "go/ast"

// callMethodName returns the method name for `<expr>.Name(...)` call
// expressions. For non-method calls (plain `foo(...)`) it returns "".
// Callers rely on the empty string to skip pattern matching on free
// functions, keeping PRV-16 focussed on receiver-style lookups.
func callMethodName(call *ast.CallExpr) string {
	if call == nil {
		return ""
	}
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok || sel.Sel == nil {
		return ""
	}
	return sel.Sel.Name
}
