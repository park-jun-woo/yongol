//ff:func feature=validate-contract type=util control=sequence topic=preserve-safety
//ff:what isErrCheckIf — *ast.IfStmt 이 err != nil 계열 가드인지 판정

package contract

import (
	"go/ast"
	"go/token"
)

// isErrCheckIf reports whether stmt is an `if <errIdent> != nil { ... }`
// style guard. A leading simple-init (`if err := foo(); err != nil`) is
// tolerated. errName lets callers pin the expected identifier; an empty
// string accepts any identifier whose name ends with "err" or "Err" so
// common shapes like `if dbErr != nil` still count.
func isErrCheckIf(stmt *ast.IfStmt, errName string) bool {
	if stmt == nil {
		return false
	}
	bin, ok := stmt.Cond.(*ast.BinaryExpr)
	if !ok || bin.Op != token.NEQ {
		return false
	}
	return binaryIsErrNotNil(bin, errName)
}
