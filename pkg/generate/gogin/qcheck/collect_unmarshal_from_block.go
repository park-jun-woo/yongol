//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what collectUnmarshalFromBlock — 단일 BlockStmt + 중첩 block 재귀 DF-01 수집

package qcheck

import (
	"go/ast"
	"go/token"
)

// collectUnmarshalFromBlock emits findings for every statement in block
// via unmarshalInStmt and then recurses into nested *ast.BlockStmt
// children so if/for/switch bodies are covered. The recursion uses
// ast.Inspect on each nested statement — a clean iteration over
// block.List at depth 1 keeps filefunc happy.
func collectUnmarshalFromBlock(block *ast.BlockStmt, targets []string, fset *token.FileSet) []DefensiveFinding {
	var findings []DefensiveFinding
	for i, stmt := range block.List {
		findings = append(findings, unmarshalInStmt(stmt, block.List, i, targets, fset)...)
		ast.Inspect(stmt, func(n ast.Node) bool {
			inner, ok := n.(*ast.BlockStmt)
			if !ok || inner == block {
				return true
			}
			findings = append(findings, collectUnmarshalFromBlock(inner, targets, fset)...)
			return false
		})
	}
	return findings
}
