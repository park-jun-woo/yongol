//ff:func feature=gen-gogin type=util control=sequence
//ff:what stmtIsErrGuard — stmt 가 `if err != nil { ... }` 계열 err 체크 가드인지 판정

package qcheck

import (
	"go/ast"
	"go/token"
)

// stmtIsErrGuard reports whether stmt is an `if <errIdent> != nil { ... }`
// guard. A leading simple-init (`if err := foo(); err != nil`) is tolerated
// because Go idiom often wraps the triggering call into the if-init itself.
// The err identifier must literally be "err" or end with "err"/"Err" so
// common shapes (dbErr, unmarshalErr) still count. Qualified receivers are
// rejected on purpose — we want direct local error gates only.
func stmtIsErrGuard(stmt ast.Stmt) bool {
	ifStmt, ok := stmt.(*ast.IfStmt)
	if !ok {
		return false
	}
	bin, ok := ifStmt.Cond.(*ast.BinaryExpr)
	if !ok || bin.Op != token.NEQ {
		return false
	}
	return binaryIsErrNilCheck(bin)
}
