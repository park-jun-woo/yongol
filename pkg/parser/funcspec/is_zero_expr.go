//ff:func feature=funcspec type=util control=selection
//ff:what isZeroExpr — zero value expression (nil / "" / 0 / empty composite) 여부 검사
package funcspec

import (
	"go/ast"
	"go/token"
)

// isZeroExpr reports whether e is a zero-value expression: nil identifier,
// empty composite literal, or zero basic literal ("", 0, 0.0, false).
func isZeroExpr(e ast.Expr) bool {
	switch v := e.(type) {
	case *ast.Ident:
		return v.Name == "nil" || v.Name == "false"
	case *ast.CompositeLit:
		return len(v.Elts) == 0
	case *ast.BasicLit:
		switch v.Kind {
		case token.INT, token.FLOAT:
			return v.Value == "0" || v.Value == "0.0"
		case token.STRING:
			return v.Value == `""` || v.Value == "``"
		}
	}
	return false
}
