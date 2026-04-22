//ff:func feature=gen-gogin type=util control=sequence
//ff:what ifInitCalls — `if err := pkg.Func(...); err != nil` 꼴 Init 에서 대상 호출 여부 판정

package qcheck

import "go/ast"

// ifInitCalls reports whether an IfStmt's Init is an assignment whose RHS
// is a pkgName.funcName call. Used to detect the idiomatic
// `if err := json.Unmarshal(...); err != nil { ... }` pattern so the
// outer call is considered already guarded.
func ifInitCalls(ifStmt *ast.IfStmt, pkgName, funcName string) bool {
	if ifStmt == nil || ifStmt.Init == nil {
		return false
	}
	assign, ok := ifStmt.Init.(*ast.AssignStmt)
	if !ok || len(assign.Rhs) != 1 {
		return false
	}
	call, ok := assign.Rhs[0].(*ast.CallExpr)
	if !ok {
		return false
	}
	return isSelectorCall(call, pkgName, funcName)
}
