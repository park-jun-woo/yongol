//ff:func feature=funcspec type=util control=sequence
//ff:what isPanicCall — 스텁 판정용 bare panic(...) ExprStmt 여부 확인
package funcspec

import "go/ast"

// isPanicCall reports whether stmt is a bare `panic(...)` expression statement.
func isPanicCall(stmt ast.Stmt) bool {
	exprStmt, ok := stmt.(*ast.ExprStmt)
	if !ok {
		return false
	}
	callExpr, ok := exprStmt.X.(*ast.CallExpr)
	if !ok {
		return false
	}
	ident, ok := callExpr.Fun.(*ast.Ident)
	return ok && ident.Name == "panic"
}
