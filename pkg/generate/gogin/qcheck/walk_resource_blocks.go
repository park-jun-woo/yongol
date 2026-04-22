//ff:func feature=gen-gogin type=util control=iteration dimension=2
//ff:what walkResourceBlocks — file AST 전역을 돌며 리소스 close 누락 findings 누적

package qcheck

import (
	"go/ast"
	"go/token"
)

// walkResourceBlocks iterates every function body in file and delegates
// per-block scanning to collectResourceFromBlock. Depth-1 loop over
// Decls satisfies filefunc A11.
func walkResourceBlocks(file *ast.File, fset *token.FileSet) []DefensiveFinding {
	var findings []DefensiveFinding
	for _, decl := range file.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Body == nil {
			continue
		}
		findings = append(findings, collectResourceFromBlock(fn.Body, fset)...)
	}
	return findings
}
