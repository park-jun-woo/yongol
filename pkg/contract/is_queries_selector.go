//ff:func feature=contract type=util control=sequence
//ff:what isQueriesSelector — SelectorExpr 의 tail name 이 "Queries" 인지 검사 (sqlc 패턴 감지)

package contract

import "go/ast"

// isQueriesSelector reports whether x is a selector whose tail name
// is "Queries" — matching `server.Queries`, `s.Queries`, and similar
// patterns produced by sqlc consumers.
func isQueriesSelector(x ast.Expr) bool {
	sel, ok := x.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	return sel.Sel.Name == "Queries"
}
