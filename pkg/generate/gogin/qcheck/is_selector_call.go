//ff:func feature=gen-gogin type=util control=sequence
//ff:what isSelectorCall — *ast.CallExpr 이 pkg.Func (둘 다 non-empty 매칭) 호출인지 판정

package qcheck

import "go/ast"

// isSelectorCall reports whether call is `x.<funcName>(...)` where x is a
// plain identifier (package or receiver name). If pkgName is non-empty the
// receiver ident must match exactly; empty pkgName accepts any ident.
// funcName must always match. Qualified receivers like `a.b.Scan` return
// false because callers use this helper to pin direct package-level calls
// (json.Unmarshal, yaml.Unmarshal) and method calls on a named local var
// (row.Scan, rows.Scan).
func isSelectorCall(call *ast.CallExpr, pkgName, funcName string) bool {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	ident, ok := sel.X.(*ast.Ident)
	if !ok {
		return false
	}
	if pkgName != "" && ident.Name != pkgName {
		return false
	}
	return sel.Sel.Name == funcName
}
