//ff:func feature=gen-gogin type=util control=selection
//ff:what maxBlockDepth — AST 노드의 최대 nesting depth 재귀 계산 (if/for/range/switch/select)

package qcheck

import "go/ast"

// maxBlockDepth walks a statement subtree and returns the deepest control
// nesting. Each if/for/range/switch/select adds 1 to the depth of its body.
// Recursion delegates per-kind walk to walkChildDepths so the switch here
// stays flat (filefunc Q1 depth ≤ 3 for selection).
func maxBlockDepth(n ast.Node, depth int) int {
	switch node := n.(type) {
	case *ast.BlockStmt:
		return maxOfStmts(node.List, depth)
	case *ast.IfStmt:
		return ifStmtDepth(node, depth)
	case *ast.ForStmt:
		return maxBlockDepth(node.Body, depth+1)
	case *ast.RangeStmt:
		return maxBlockDepth(node.Body, depth+1)
	case *ast.SwitchStmt:
		return maxBlockDepth(node.Body, depth+1)
	case *ast.TypeSwitchStmt:
		return maxBlockDepth(node.Body, depth+1)
	case *ast.SelectStmt:
		return maxBlockDepth(node.Body, depth+1)
	case *ast.CaseClause:
		return maxOfStmts(node.Body, depth)
	case *ast.CommClause:
		return maxOfStmts(node.Body, depth)
	}
	return depth
}
