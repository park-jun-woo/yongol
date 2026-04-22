//ff:func feature=validate-contract type=util control=selection topic=preserve-safety
//ff:what binaryIsErrNotNil — `<ident> != nil` 혹은 `nil != <ident>` 패턴 판정

package contract

import (
	"go/ast"
	"strings"
)

// binaryIsErrNotNil reports whether bin is `x != nil` or `nil != x`
// where x is either errName (if non-empty) or any identifier whose
// textual name ends with "err"/"Err". Only plain idents are considered
// — qualified selectors (`foo.err`) are intentionally excluded because
// the caller uses this to detect direct local error gates.
func binaryIsErrNotNil(bin *ast.BinaryExpr, errName string) bool {
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
	}
	if errIdent == nil {
		return false
	}
	if errName != "" {
		return errIdent.Name == errName
	}
	lower := strings.ToLower(errIdent.Name)
	return lower == "err" || strings.HasSuffix(lower, "err")
}
