//ff:func feature=gen-gogin type=util control=iteration dimension=2
//ff:what walkUnmarshalBlocks — file AST 전역을 돌며 unmarshal 미가드 findings 누적

package qcheck

import (
	"go/ast"
	"go/token"
)

// walkUnmarshalBlocks traverses file's AST block-by-block and calls
// unmarshalInStmt on every statement to collect DF-01 findings. The
// traversal uses ast.Inspect internally but exposes an explicit loop
// over block.List so the file satisfies filefunc A11 (control=iteration
// must have a loop at depth 1).
func walkUnmarshalBlocks(file *ast.File, fset *token.FileSet) []DefensiveFinding {
	targets := []string{"json", "yaml"}
	var findings []DefensiveFinding
	for _, decl := range file.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Body == nil {
			continue
		}
		findings = append(findings, collectUnmarshalFromBlock(fn.Body, targets, fset)...)
	}
	return findings
}
