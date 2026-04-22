//ff:func feature=validate-contract type=util control=selection topic=preserve-safety
//ff:what unmarshalAssignInStmt — stmt 가 Unmarshal 호출이면 assign/call/err 이름/분류 반환

package contract

import "go/ast"

// unmarshalAssignInStmt inspects stmt for a json/yaml Unmarshal call
// and reports what shape it takes. Possible outcomes:
//
//   - Not an Unmarshal statement → (nil, nil, "", unknown).
//   - `if err := pkg.Unmarshal(...); err != nil { ... }` → Discarded
//     (the if condition itself is the guard).
//   - `_ = pkg.Unmarshal(...)` → Discarded.
//   - `err := pkg.Unmarshal(...)` → Assigned with errName="err" for
//     hasErrCheckAfter to follow up.
//   - `pkg.Unmarshal(...)` as ExprStmt → Assigned with errName="" so
//     the caller raises PRV-12 immediately (no receiver to guard).
func unmarshalAssignInStmt(stmt ast.Stmt) (*ast.AssignStmt, *ast.CallExpr, string, unmarshalKind) {
	switch s := stmt.(type) {
	case *ast.IfStmt:
		return unmarshalInIfStmt(s)
	case *ast.AssignStmt:
		return unmarshalInAssignStmt(s)
	case *ast.ExprStmt:
		return unmarshalInExprStmt(s)
	}
	return nil, nil, "", unmarshalKindUnknown
}
