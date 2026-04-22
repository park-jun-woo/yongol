//ff:func feature=validate-contract type=util control=iteration dimension=1 topic=preserve-safety
//ff:what hasErrCheckAfter — 블록 내 stmtIdx 의 호출 결과 err 를 뒤따르는 stmt 가 체크하는지

package contract

import "go/ast"

// hasErrCheckAfter scans stmts[stmtIdx+1:] for an `if <errName> != nil`
// guard. It stops searching once an assignment that rebinds errName to
// something that is NOT another tracked call appears — that is the
// point where the error value has been clobbered and any later check
// refers to a different value.
//
// errName is the identifier that carries the error returned by the
// statement at stmtIdx. Pass "" to accept any identifier whose name is
// `err` or ends with `Err`/`err` (common shapes).
//
// A returned true means: "The error surfaced by stmtIdx is properly
// guarded somewhere later in the same block". A returned false means
// the rule should emit a diagnostic.
func hasErrCheckAfter(stmts []ast.Stmt, stmtIdx int, errName string) bool {
	if stmtIdx < 0 || stmtIdx >= len(stmts) {
		return false
	}
	for i := stmtIdx + 1; i < len(stmts); i++ {
		if ifs, ok := stmts[i].(*ast.IfStmt); ok && isErrCheckIf(ifs, errName) {
			return true
		}
		if assignClobbersErr(stmts[i], errName) {
			return false
		}
	}
	return false
}
