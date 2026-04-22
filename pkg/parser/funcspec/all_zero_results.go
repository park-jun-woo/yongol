//ff:func feature=funcspec type=util control=iteration dimension=1
//ff:what allZeroResults — return 결과 expr 들이 모두 zero value 인지 검사
package funcspec

import "go/ast"

// allZeroResults reports whether every expression is a zero value (nil
// identifier, empty composite literal, or zero basic literal).
// An empty result list (bare `return` in a void func) is also zero.
func allZeroResults(results []ast.Expr) bool {
	for _, e := range results {
		if !isZeroExpr(e) {
			return false
		}
	}
	return true
}
