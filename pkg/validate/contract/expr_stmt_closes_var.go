//ff:func feature=validate-contract type=util control=sequence topic=preserve-safety
//ff:what exprStmtClosesVar — ExprStmt 가 varName 에 대한 Close() 호출 문장인지

package contract

import "go/ast"

// exprStmtClosesVar reports whether stmt is `varName.Close()` as a
// bare expression statement (the common shape inside a deferred
// closure body). Assignment-form close calls (`err = varName.Close()`)
// are also accepted — the outer scanners care only that Close is
// invoked for varName before the function returns.
func exprStmtClosesVar(stmt ast.Stmt, varName string) bool {
	if es, ok := stmt.(*ast.ExprStmt); ok {
		if call, ok := es.X.(*ast.CallExpr); ok {
			return callClosesVar(call, varName)
		}
	}
	if as, ok := stmt.(*ast.AssignStmt); ok && len(as.Rhs) == 1 {
		if call, ok := as.Rhs[0].(*ast.CallExpr); ok {
			return callClosesVar(call, varName)
		}
	}
	return false
}
