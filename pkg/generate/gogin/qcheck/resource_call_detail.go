//ff:func feature=gen-gogin type=util control=selection
//ff:what resourceCallDetail — CallExpr 이 리소스 획득 호출이면 "pkg.Func" 문자열 반환

package qcheck

import "go/ast"

// resourceCallDetail returns a "pkg.Func" label when call matches one of
// the known resource-returning forms (os.Open, <x>.Query, <x>.QueryContext,
// <x>.Prepare, <x>.PrepareContext). Returning an empty string signals the
// caller to skip this statement. DB method receivers vary (db, conn, tx,
// qtx) so pkgName is not pinned — the label uses whatever ident was seen.
func resourceCallDetail(call *ast.CallExpr) string {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return ""
	}
	ident, ok := sel.X.(*ast.Ident)
	if !ok {
		return ""
	}
	name := sel.Sel.Name
	if ident.Name == "os" && name == "Open" {
		return "os.Open"
	}
	switch name {
	case "Query", "QueryContext", "Prepare", "PrepareContext":
		return ident.Name + "." + name
	}
	return ""
}
