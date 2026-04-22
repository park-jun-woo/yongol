//ff:func feature=validate-contract type=util control=sequence topic=preserve-safety
//ff:what unmarshalInExprStmt — ExprStmt 가 Unmarshal 호출이면 Assigned 로 분류

package contract

import "go/ast"

// unmarshalInExprStmt handles the bare-call shape
// `json.Unmarshal(body, req)` that happens when a preserved edit
// dropped the error variable entirely. There is no receiver to check,
// so the caller (unmarshalDiagsInBlock) emits PRV-12 unconditionally.
func unmarshalInExprStmt(s *ast.ExprStmt) (*ast.AssignStmt, *ast.CallExpr, string, unmarshalKind) {
	call, ok := s.X.(*ast.CallExpr)
	if !ok || !isUnmarshalCall(call) {
		return nil, nil, "", unmarshalKindUnknown
	}
	return nil, call, "", unmarshalKindAssigned
}
