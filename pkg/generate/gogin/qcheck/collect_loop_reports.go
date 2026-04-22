//ff:func feature=gen-gogin type=util control=sequence
//ff:what collectLoopReports — FuncDecl 내 모든 for/range 루프를 PureLinesReport로 변환

package qcheck

import (
	"go/ast"
	"go/token"
)

// collectLoopReports walks a FuncDecl and emits one PureLinesReport per
// for/range loop, using source offsets to count non-blank non-comment lines
// inside each loop body. Per-loop body line counting delegates to bodyReport
// so this file stays 1-func.
func collectLoopReports(fset *token.FileSet, fn *ast.FuncDecl, src string) []PureLinesReport {
	var out []PureLinesReport
	ast.Inspect(fn, func(node ast.Node) bool {
		appendLoopReport(&out, fset, fn.Name.Name, node, src)
		return true
	})
	return out
}
