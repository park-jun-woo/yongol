//ff:func feature=validate-contract type=util control=sequence topic=preserve-safety
//ff:what unmarshalInIfStmt — IfStmt.Init 이 Unmarshal 호출이면 Discarded 로 반환

package contract

import "go/ast"

// unmarshalInIfStmt recognises `if err := json.Unmarshal(...); err != nil`
// — the canonical safe shape. The call is reported as Discarded
// because the if condition necessarily guards the error, even if the
// body itself is empty (Go rejects unused err init otherwise).
func unmarshalInIfStmt(s *ast.IfStmt) (*ast.AssignStmt, *ast.CallExpr, string, unmarshalKind) {
	as, ok := s.Init.(*ast.AssignStmt)
	if !ok {
		return nil, nil, "", unmarshalKindUnknown
	}
	call := unmarshalCallFromAssign(as)
	if call == nil {
		return nil, nil, "", unmarshalKindUnknown
	}
	return as, call, "", unmarshalKindDiscarded
}
