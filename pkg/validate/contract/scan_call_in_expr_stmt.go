//ff:func feature=validate-contract type=util control=sequence topic=preserve-safety
//ff:what scanCallInExprStmt — ExprStmt 가 Scan 호출이면 kind=Assigned 로 분류

package contract

import "go/ast"

// scanCallInExprStmt handles the bare-call shape `row.Scan(&x)` — no
// error receiver, so it is immediately Assigned with an empty errName.
// The caller (scanDiagsInBlock) will raise PRV-13 because there is no
// way to guard an ignored error return.
func scanCallInExprStmt(s *ast.ExprStmt) (*ast.CallExpr, string, scanKind) {
	call, ok := s.X.(*ast.CallExpr)
	if !ok || !isScanCall(call) {
		return nil, "", scanKindUnknown
	}
	return call, "", scanKindAssigned
}
