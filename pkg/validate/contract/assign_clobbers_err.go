//ff:func feature=validate-contract type=util control=sequence topic=preserve-safety
//ff:what assignClobbersErr — stmt 가 errName 을 새 값으로 덮어쓰는 assignment 인지 판정

package contract

import "go/ast"

// assignClobbersErr reports whether stmt is an AssignStmt that rebinds
// errName to a fresh value, destroying the error produced by the
// tracked call. This is used by hasErrCheckAfter to stop scanning once
// the original error value has been overwritten — a later `if err !=
// nil` would then refer to the NEW error, not the one under audit.
//
// A bare `err = nil` is treated as an EXPLICIT discard (not clobber):
// the preserved file author is acknowledging they do not care about
// the value, which is different from silently ignoring it.
func assignClobbersErr(stmt ast.Stmt, errName string) bool {
	assign, ok := stmt.(*ast.AssignStmt)
	if !ok || len(assign.Lhs) == 0 {
		return false
	}
	idx := indexOfErrLhs(assign, errName)
	if idx < 0 {
		return false
	}
	if rhsAtIndexIsNil(assign, idx) {
		return false
	}
	return true
}
