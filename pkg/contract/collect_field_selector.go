//ff:func feature=contract type=util control=sequence
//ff:what collectFieldSelector — CallExpr 바깥 SelectorExpr 중 exported 필드 접근을 fields 맵에 적재

package contract

import (
	"go/ast"
	"go/token"
)

// collectFieldSelector is the ast.Inspect visitor for the second
// pass: it collects SelectorExpr nodes that are not call targets
// and whose selector name starts with an upper-case letter. Only
// simple Ident receivers produce entries — chained selectors and
// call/index expressions are skipped because they rarely map cleanly
// to a DDL column without significantly more analysis.
func collectFieldSelector(
	fset *token.FileSet,
	n ast.Node,
	callSelectors map[*ast.SelectorExpr]struct{},
	fields map[string]struct{},
) bool {
	sel, ok := n.(*ast.SelectorExpr)
	if !ok {
		return true
	}
	if _, isCall := callSelectors[sel]; isCall {
		return true
	}
	if !isExported(sel.Sel.Name) {
		return true
	}
	recv := renderRecv(fset, sel.X)
	if recv == "" {
		return true
	}
	fields[recv+"."+sel.Sel.Name] = struct{}{}
	return true
}
