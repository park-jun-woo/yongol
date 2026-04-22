//ff:func feature=gen-splitter type=util control=selection
//ff:what receiverName — 메서드 receiver 표현식에서 타입 이름 추출 (T / *T 지원)
package splitter

import "go/ast"

// receiverName returns the named type of a method receiver expression.
// It handles the common forms "T" and "*T"; other expressions (generic
// instantiations, etc.) yield an empty string so callers can fall back
// to a default file name.
func receiverName(expr ast.Expr) string {
	switch e := expr.(type) {
	case *ast.Ident:
		return e.Name
	case *ast.StarExpr:
		if id, ok := e.X.(*ast.Ident); ok {
			return id.Name
		}
	}
	return ""
}
