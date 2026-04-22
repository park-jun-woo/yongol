//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what collectDepthReports — 파일의 모든 top-level FuncDecl → DepthReport 변환

package qcheck

import "go/ast"

// collectDepthReports walks top-level declarations and emits a DepthReport
// for each FuncDecl with a body. Non-function declarations are skipped.
func collectDepthReports(file *ast.File) []DepthReport {
	var reports []DepthReport
	for _, decl := range file.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Body == nil {
			continue
		}
		reports = append(reports, DepthReport{
			Func:     fn.Name.Name,
			MaxDepth: maxBlockDepth(fn.Body, 0),
		})
	}
	return reports
}
