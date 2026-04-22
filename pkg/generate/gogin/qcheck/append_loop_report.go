//ff:func feature=gen-gogin type=util control=selection
//ff:what appendLoopReport — AST 노드가 for/range 이면 PureLinesReport 하나 append

package qcheck

import (
	"go/ast"
	"go/token"
)

// appendLoopReport checks whether node is a loop and, if so, appends a
// PureLinesReport to out. Non-loop nodes are ignored. Splitting this from
// collectLoopReports keeps both functions flat (filefunc Q1 / Q4).
func appendLoopReport(out *[]PureLinesReport, fset *token.FileSet, fnName string, node ast.Node, src string) {
	switch loop := node.(type) {
	case *ast.ForStmt:
		*out = append(*out, bodyReport(fset, fnName, "for", loop.For, loop.Body, src))
	case *ast.RangeStmt:
		*out = append(*out, bodyReport(fset, fnName, "range", loop.For, loop.Body, src))
	}
}
