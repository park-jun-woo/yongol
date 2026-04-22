//ff:func feature=validate-contract type=util control=sequence topic=preserve-safety
//ff:what unmarshalInAssignStmt — AssignStmt 가 Unmarshal 호출이면 kind/errName 분류

package contract

import "go/ast"

// unmarshalInAssignStmt classifies the `err := json.Unmarshal(...)`
// / `_ = json.Unmarshal(...)` shapes. Blank-discard LHS maps to
// Discarded; any other ident is Assigned with unmarshalErrName
// carrying forward the tracked identifier.
func unmarshalInAssignStmt(s *ast.AssignStmt) (*ast.AssignStmt, *ast.CallExpr, string, unmarshalKind) {
	call := unmarshalCallFromAssign(s)
	if call == nil {
		return nil, nil, "", unmarshalKindUnknown
	}
	if assignIsBlankDiscard(s) {
		return s, call, "", unmarshalKindDiscarded
	}
	return s, call, unmarshalErrName(s), unmarshalKindAssigned
}
