//ff:func feature=gen-gogin type=util control=sequence
//ff:what ifStmtDepth — IfStmt 의 body + else 중 최대 depth 계산

package qcheck

import "go/ast"

// ifStmtDepth returns the deeper of the IfStmt body and its else branch,
// each measured at depth+1 because entering the if adds one nesting level.
// Extracted from maxBlockDepth to keep that switch body within filefunc Q4.
func ifStmtDepth(node *ast.IfStmt, depth int) int {
	max := maxBlockDepth(node.Body, depth+1)
	if node.Else == nil {
		return max
	}
	if child := maxBlockDepth(node.Else, depth+1); child > max {
		max = child
	}
	return max
}
