//ff:func feature=ssac-parse type=util control=selection
//ff:what AST 표현식을 문자열로 변환
package ssac

import "go/ast"

// exprToString은 AST 표현식을 문자열로 변환한다.
func exprToString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return exprToString(t.X) + "." + t.Sel.Name
	default:
		return ""
	}
}
