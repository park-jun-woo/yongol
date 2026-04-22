//ff:func feature=gen-gogin type=util control=iteration dimension=2
//ff:what walkScanBlocks — file AST 전역을 돌며 `.Scan` 미가드 findings 누적

package qcheck

import (
	"go/ast"
	"go/token"
)

// walkScanBlocks iterates every function body in file and asks
// collectScanFromBlock to produce DF-02 findings for it. Keeping the
// loop at depth 1 (over Decls) satisfies filefunc A11.
func walkScanBlocks(file *ast.File, fset *token.FileSet) []DefensiveFinding {
	var findings []DefensiveFinding
	for _, decl := range file.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Body == nil {
			continue
		}
		findings = append(findings, collectScanFromBlock(fn.Body, fset)...)
	}
	return findings
}
