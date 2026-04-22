//ff:func feature=gen-gogin type=util control=selection
//ff:what binaryIsErrNilCheck — `<errIdent> != nil` 혹은 `nil != <errIdent>` 패턴 판정

package qcheck

import (
	"go/ast"
	"strings"
)

// binaryIsErrNilCheck reports whether bin is `<x> != nil` or `nil != <x>`
// where x is a plain Ident whose lowercased name equals "err" or ends
// with "err". The switch keeps the selection at depth 1 (filefunc A10).
func binaryIsErrNilCheck(bin *ast.BinaryExpr) bool {
	lhs, lok := bin.X.(*ast.Ident)
	rhs, rok := bin.Y.(*ast.Ident)
	if !lok || !rok {
		return false
	}
	var errIdent *ast.Ident
	switch {
	case lhs.Name == "nil" && rhs.Name != "nil":
		errIdent = rhs
	case rhs.Name == "nil" && lhs.Name != "nil":
		errIdent = lhs
	default:
		return false
	}
	lower := strings.ToLower(errIdent.Name)
	return lower == "err" || strings.HasSuffix(lower, "err")
}
