//ff:func feature=validate-contract type=util control=sequence topic=preserve-safety
//ff:what scanCallInAssignStmt — AssignStmt 가 Scan 호출이면 kind/errName 반환

package contract

import "go/ast"

// scanCallInAssignStmt classifies the `v := row.Scan(...)` /
// `_ = row.Scan(...)` shapes. Blank-discard LHS maps to Discarded;
// any other ident name goes through unmarshalErrName (reused because
// the heuristic is identical — any name ending in "err" qualifies).
func scanCallInAssignStmt(s *ast.AssignStmt) (*ast.CallExpr, string, scanKind) {
	call := scanCallFromAssign(s)
	if call == nil {
		return nil, "", scanKindUnknown
	}
	if assignIsBlankDiscard(s) {
		return call, "", scanKindDiscarded
	}
	return call, unmarshalErrName(s), scanKindAssigned
}
