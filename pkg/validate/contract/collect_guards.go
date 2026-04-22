//ff:func feature=validate-contract type=util control=sequence topic=preserve-safety
//ff:what collectGuards — 함수 body 내 len(x)/range x 언급을 guard map 에 기록

package contract

import "go/ast"

// collectGuards walks body and marks every identifier that appears as
// the argument of `len(x)` OR the target of a `for ... := range x`
// loop header as guarded. Subsequent sliceBoundsDiagsInFunc consults
// this map to skip indexing expressions on names the author has
// already surveyed. The walk is eager — we do not track ordering,
// because preserved handlers are small enough that any guard in the
// same function is evidence of intent.
func collectGuards(body *ast.BlockStmt, guarded map[string]bool) {
	ast.Inspect(body, func(n ast.Node) bool {
		if call, ok := n.(*ast.CallExpr); ok {
			if markLenGuard(call, guarded) {
				return true
			}
		}
		if rs, ok := n.(*ast.RangeStmt); ok && rs.X != nil {
			if name := leftmostIdentName(rs.X); name != "" {
				guarded[name] = true
			}
		}
		return true
	})
}
