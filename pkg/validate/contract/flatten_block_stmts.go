//ff:func feature=validate-contract type=util control=sequence topic=preserve-safety
//ff:what flattenBlockStmts — BlockStmt 트리에서 모든 하위 stmt 를 단일 슬라이스로 수집

package contract

import "go/ast"

// flattenBlockStmts returns every Stmt reachable from body — across
// nested blocks, if/else arms, and for loops — as a single flat
// slice. hasDeferClose uses this to discover defers that appear in
// sibling branches (legitimate) while still being cheap to compute.
//
// FuncLit bodies are intentionally NOT descended into: a defer inside
// a closure runs in that closure's scope, not the enclosing function,
// so it cannot satisfy the outer resource-close contract.
func flattenBlockStmts(body *ast.BlockStmt) []ast.Stmt {
	var out []ast.Stmt
	ast.Inspect(body, func(n ast.Node) bool {
		if _, ok := n.(*ast.FuncLit); ok {
			return false
		}
		if s, ok := n.(ast.Stmt); ok {
			out = append(out, s)
		}
		return true
	})
	return out
}
