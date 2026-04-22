//ff:func feature=validate-contract type=util control=sequence topic=preserve-safety
//ff:what isZeroIndex — expr 가 리터럴 `0` 인지 판정 (PRV-14 전용)

package contract

import (
	"go/ast"
	"go/token"
)

// isZeroIndex reports whether expr is the integer literal "0". Only
// this literal shape is interesting to PRV-14 — variable-backed
// indexing is too ambiguous to flag without type-level reasoning.
func isZeroIndex(expr ast.Expr) bool {
	lit, ok := expr.(*ast.BasicLit)
	if !ok {
		return false
	}
	return lit.Kind == token.INT && lit.Value == "0"
}
