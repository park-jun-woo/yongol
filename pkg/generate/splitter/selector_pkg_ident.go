//ff:func feature=gen-splitter type=util control=sequence
//ff:what selectorPkgIdent — ast.Node 가 SelectorExpr 인 경우 X(pkg 식별자) 반환, 아니면 nil
package splitter

import "go/ast"

// selectorPkgIdent returns the leading identifier of a SelectorExpr
// when n is of that shape and its X side is a simple *ast.Ident (the
// canonical pkg.Name pattern). Method calls on non-ident receivers
// (chained selectors, function calls) yield nil, which gather callers
// treat as "not an import reference".
func selectorPkgIdent(n ast.Node) *ast.Ident {
	sel, ok := n.(*ast.SelectorExpr)
	if !ok {
		return nil
	}
	id, ok := sel.X.(*ast.Ident)
	if !ok {
		return nil
	}
	return id
}
