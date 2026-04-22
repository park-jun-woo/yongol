//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what maxOfStmts — 여러 stmt 중 최대 nesting depth 반환

package qcheck

import "go/ast"

// maxOfStmts returns the deepest depth across a slice of statements, each
// measured at the caller's current depth. Used for BlockStmt / CaseClause /
// CommClause bodies which hold a list of statements but don't themselves
// add a nesting level.
func maxOfStmts(stmts []ast.Stmt, depth int) int {
	max := depth
	for _, stmt := range stmts {
		if child := maxBlockDepth(stmt, depth); child > max {
			max = child
		}
	}
	return max
}
