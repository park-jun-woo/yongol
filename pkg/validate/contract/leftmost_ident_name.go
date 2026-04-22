//ff:func feature=validate-contract type=util control=iteration dimension=1 topic=preserve-safety
//ff:what leftmostIdentName — 중첩 SelectorExpr 의 가장 왼쪽 Ident 이름 반환

package contract

import "go/ast"

// leftmostIdentName walks nested SelectorExpr nodes leftward until it
// reaches an *ast.Ident and returns its name. Non-ident roots (calls,
// literals, type assertions) yield "". Used by call-scanners that want
// to identify the "owning" variable of a chained expression such as
// `resp.Body.Close()` (leftmost = "resp").
func leftmostIdentName(expr ast.Expr) string {
	for {
		switch v := expr.(type) {
		case *ast.Ident:
			return v.Name
		case *ast.SelectorExpr:
			expr = v.X
		default:
			return ""
		}
	}
}
