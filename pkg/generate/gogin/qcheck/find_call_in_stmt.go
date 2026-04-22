//ff:func feature=gen-gogin type=util control=sequence
//ff:what findCallInStmt — stmt 안에서 pkg.funcName 호출 *ast.CallExpr 를 찾아 반환

package qcheck

import "go/ast"

// findCallInStmt walks stmt and returns the first CallExpr matching
// pkgName.funcName (direct selector form). Used by scanners to know
// whether a statement textually contains the target call regardless of
// AssignStmt / IfStmt.Init / ExprStmt shape. Returns nil when absent.
func findCallInStmt(stmt ast.Stmt, pkgName, funcName string) *ast.CallExpr {
	var found *ast.CallExpr
	ast.Inspect(stmt, func(n ast.Node) bool {
		if found != nil {
			return false
		}
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}
		if isSelectorCall(call, pkgName, funcName) {
			found = call
			return false
		}
		return true
	})
	return found
}
