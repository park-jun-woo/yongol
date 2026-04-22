//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what collectResourceFromBlock — 단일 BlockStmt + 중첩 block 재귀 DF-06 수집

package qcheck

import (
	"go/ast"
	"go/token"
)

// collectResourceFromBlock iterates block statements for DF-06 via
// resourceCallInStmt and recurses into nested *ast.BlockStmt children
// so resource acquisitions inside if/for bodies are also covered. The
// recursion uses ast.Inspect over nested statements, keeping the explicit
// loop at depth 1.
func collectResourceFromBlock(block *ast.BlockStmt, fset *token.FileSet) []DefensiveFinding {
	var findings []DefensiveFinding
	for i, stmt := range block.List {
		findings = append(findings, resourceCallInStmt(stmt, block.List, i, fset)...)
		ast.Inspect(stmt, func(n ast.Node) bool {
			inner, ok := n.(*ast.BlockStmt)
			if !ok || inner == block {
				return true
			}
			findings = append(findings, collectResourceFromBlock(inner, fset)...)
			return false
		})
	}
	return findings
}
