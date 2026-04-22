//ff:func feature=gen-gogin type=util control=sequence
//ff:what resourceCallInStmt — 한 Stmt 에서 리소스 획득 호출을 찾아 defer Close 후속 여부 판정

package qcheck

import (
	"go/ast"
	"go/token"
)

// resourceCallInStmt checks whether stmt is an AssignStmt whose RHS is a
// known resource-returning call (db.Query*, os.Open). When such a call is
// found, it scans blockList[i+1:] for a matching `defer <recv>.Close()`.
// Absence yields a DefensiveFinding. Non-assign statements and statements
// whose RHS isn't a resource call are silently skipped.
func resourceCallInStmt(stmt ast.Stmt, blockList []ast.Stmt, i int, fset *token.FileSet) []DefensiveFinding {
	assign, ok := stmt.(*ast.AssignStmt)
	if !ok || len(assign.Rhs) != 1 || len(assign.Lhs) == 0 {
		return nil
	}
	call, ok := assign.Rhs[0].(*ast.CallExpr)
	if !ok {
		return nil
	}
	detail := resourceCallDetail(call)
	if detail == "" {
		return nil
	}
	recv, ok := assign.Lhs[0].(*ast.Ident)
	if !ok || recv.Name == "_" {
		return nil
	}
	if hasDeferCloseAfter(blockList, i, recv.Name) {
		return nil
	}
	return []DefensiveFinding{{
		Category: "DF-06",
		Detail:   detail,
		Pos:      fset.Position(call.Pos()),
	}}
}
