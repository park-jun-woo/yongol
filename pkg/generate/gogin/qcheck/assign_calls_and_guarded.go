//ff:func feature=gen-gogin type=util control=sequence
//ff:what assignCallsAndGuarded — AssignStmt RHS 가 pkg.Func 이고 다음 Stmt 가 err 가드인지 판정

package qcheck

import "go/ast"

// assignCallsAndGuarded reports whether assign's RHS is a single
// pkgName.funcName call AND the statement immediately after (blockList[i+1])
// is an `if err != nil { ... }` guard. Used to accept the two-line shape
//
//	err := pkg.Func(...)
//	if err != nil { return err }
//
// as properly guarded. Assignments whose LHS drops the error with `_` are
// intentionally not covered here — the caller's scanner still reports
// those via findCallInStmt fallback when guarded=false.
func assignCallsAndGuarded(assign *ast.AssignStmt, pkgName, funcName string, blockList []ast.Stmt, i int) bool {
	if len(assign.Rhs) != 1 {
		return false
	}
	call, ok := assign.Rhs[0].(*ast.CallExpr)
	if !ok {
		return false
	}
	if !isSelectorCall(call, pkgName, funcName) {
		return false
	}
	if !assignLHSHasErr(assign) {
		return false
	}
	if i+1 >= len(blockList) {
		return false
	}
	return stmtIsErrGuard(blockList[i+1])
}
