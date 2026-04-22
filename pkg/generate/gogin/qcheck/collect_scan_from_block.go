//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what collectScanFromBlock — 단일 BlockStmt + 중첩 block 재귀 DF-02 수집

package qcheck

import (
	"go/ast"
	"go/token"
)

// collectScanFromBlock mirrors collectUnmarshalFromBlock but calls
// scanCallInStmt for DF-02. Nested blocks (inside if/for/switch) are
// traversed via ast.Inspect so findings at any depth are captured.
func collectScanFromBlock(block *ast.BlockStmt, fset *token.FileSet) []DefensiveFinding {
	var findings []DefensiveFinding
	for i, stmt := range block.List {
		findings = append(findings, scanCallInStmt(stmt, block.List, i, fset)...)
		ast.Inspect(stmt, func(n ast.Node) bool {
			inner, ok := n.(*ast.BlockStmt)
			if !ok || inner == block {
				return true
			}
			findings = append(findings, collectScanFromBlock(inner, fset)...)
			return false
		})
	}
	return findings
}
