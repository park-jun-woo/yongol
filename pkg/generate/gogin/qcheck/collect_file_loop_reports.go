//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what collectFileLoopReports — 파일의 모든 FuncDecl 에서 루프 리포트 집계

package qcheck

import (
	"go/ast"
	"go/token"
)

// collectFileLoopReports iterates top-level FuncDecls and concatenates per-
// func loop reports. Skips non-FuncDecls and body-less decls.
func collectFileLoopReports(fset *token.FileSet, file *ast.File, src string) []PureLinesReport {
	var reports []PureLinesReport
	for _, decl := range file.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Body == nil {
			continue
		}
		reports = append(reports, collectLoopReports(fset, fn, src)...)
	}
	return reports
}
