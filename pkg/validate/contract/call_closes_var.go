//ff:func feature=validate-contract type=util control=sequence topic=preserve-safety
//ff:what callClosesVar — *ast.CallExpr 가 varName 에 대한 `.Close()` 호출인지 판정

package contract

import "go/ast"

// callClosesVar returns true when call is `<something>.Close()` and the
// leftmost identifier in the selector chain matches varName. Chained
// selectors (`varName.Body.Close()`) are accepted — if varName is
// non-nil the sub-selector is always safe to close in the stdlib
// idioms we target (http.Response.Body, etc).
func callClosesVar(call *ast.CallExpr, varName string) bool {
	if call == nil {
		return false
	}
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok || sel.Sel == nil || sel.Sel.Name != "Close" {
		return false
	}
	return leftmostIdentName(sel.X) == varName
}
