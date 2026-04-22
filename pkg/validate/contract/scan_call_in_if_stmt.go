//ff:func feature=validate-contract type=util control=sequence topic=preserve-safety
//ff:what scanCallInIfStmt — IfStmt.Init 이 Scan 호출이면 CallExpr 반환

package contract

import "go/ast"

// scanCallInIfStmt recognises the `if err := row.Scan(...); err != nil`
// init-form. The Scan call counts as handled (Discarded) because the
// if condition itself is the guard — whatever the user puts in the
// body, the returned error cannot silently escape.
func scanCallInIfStmt(s *ast.IfStmt) (*ast.CallExpr, string, scanKind) {
	as, ok := s.Init.(*ast.AssignStmt)
	if !ok {
		return nil, "", scanKindUnknown
	}
	call := scanCallFromAssign(as)
	if call == nil {
		return nil, "", scanKindUnknown
	}
	return call, "", scanKindDiscarded
}
