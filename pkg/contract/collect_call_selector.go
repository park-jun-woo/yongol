//ff:func feature=contract type=util control=sequence
//ff:what collectCallSelector — CallExpr 의 Fun SelectorExpr 을 queries/calls 맵에 분류 적재

package contract

import "go/ast"

// collectCallSelector is the ast.Inspect visitor for the first pass:
// it identifies CallExpr whose function is a SelectorExpr, tags that
// selector in callSelectors (so the second pass can skip it), and
// routes the call to either the queries map (methods on a `.Queries`
// receiver) or the calls map (package-qualified exported calls).
func collectCallSelector(
	n ast.Node,
	queries map[string]struct{},
	calls map[string]struct{},
	callSelectors map[*ast.SelectorExpr]struct{},
) bool {
	call, ok := n.(*ast.CallExpr)
	if !ok {
		return true
	}
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return true
	}
	callSelectors[sel] = struct{}{}
	if isQueriesSelector(sel.X) {
		queries[sel.Sel.Name] = struct{}{}
		return true
	}
	ident, ok := sel.X.(*ast.Ident)
	if !ok {
		return true
	}
	if _, denied := callExprPkgDenylist[ident.Name]; denied {
		return true
	}
	if isExported(sel.Sel.Name) {
		calls[ident.Name+"."+sel.Sel.Name] = struct{}{}
	}
	return true
}
